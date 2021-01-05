package main

import (
	"context"
	"fmt"
	"iloMusiAko/ent"
	"iloMusiAko/ent/user"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/diamondburned/arikawa/discord"
	"github.com/diamondburned/arikawa/gateway"
)

type state int

const (
	sending state = iota
	voting
)

var gameStates = map[discord.ChannelID]state{}
var voteStates = map[discord.ChannelID]voteData{}
var dataStates = map[discord.ChannelID]sendData{}
var ctx = context.Background()

type voteData struct {
	votes    map[discord.UserID]int
	hasVoted map[discord.UserID]struct{}
	keys     map[int]discord.UserID
}
type sendData struct {
	letters []string
	phrases map[discord.UserID][]string
}

const duration = 70

func votePhase(c *gateway.MessageCreateEvent) {
	if c.WebhookID.IsValid() {
		return
	}

	if state, ok := gameStates[c.ChannelID]; ok && state == voting {
	} else {
		return
	}

	if _, ok := voteStates[c.ChannelID].hasVoted[c.Author.ID]; ok {
		return
	}

	data, err := strconv.ParseInt(c.Content, 10, 64)
	if err != nil {
		return
	}

	id, ok := voteStates[c.ChannelID].keys[int(data)]
	if !ok {
		return
	}

	count, _ := voteStates[c.ChannelID].votes[id]

	voteStates[c.ChannelID].votes[id] = count + 1
	voteStates[c.ChannelID].hasVoted[c.Author.ID] = struct{}{}

	bot.DeleteMessage(c.ChannelID, c.ID)
	bot.SendMessage(c.ChannelID, fmt.Sprintf("<@%d> li toki e wile ona!", c.Author.ID), nil)
}

var punctuationStripper = strings.NewReplacer(
	",", "",
	".", "",
	";", "",
	":", "",
	"!", "",
	".", "",
)

func dataPhase(c *gateway.MessageCreateEvent) {
	if c.WebhookID.IsValid() {
		return
	}

	if state, ok := gameStates[c.ChannelID]; ok && state == sending {
	} else {
		return
	}

	fields := strings.Fields(c.Content)
	fieldsWithoutParticles := []string{}

	for _, field := range fields {
		if _, ok := particles[punctuationStripper.Replace(field)]; ok {
			continue
		}
		fieldsWithoutParticles = append(fieldsWithoutParticles, field)
	}

	if len(fieldsWithoutParticles) != len(dataStates[c.ChannelID].letters) {
		return
	}

	for i := 0; i < len(fieldsWithoutParticles); i++ {
		fieldsFirst := fieldsWithoutParticles[i][0:1]
		letter := dataStates[c.ChannelID].letters[i]

		if strings.ToLower(fieldsFirst) != letter {
			return
		}
	}

	dataStates[c.ChannelID].phrases[c.Author.ID] = fields
	bot.DeleteMessage(c.ChannelID, c.ID)
}

func startGame(c *gateway.MessageCreateEvent) {
	if c.WebhookID.IsValid() {
		return
	}

	rand.Seed(c.ID.Time().Unix())

	if !strings.HasPrefix(c.Content, "ilo o open e musi Ako") {
		return
	}

	if _, ok := gameStates[c.ChannelID]; ok {
		bot.SendMessage(c.ChannelID, "musi li lon. o musi kepeken ona!", nil)
		return
	}

	data := randomLetters()

	bot.SendMessage(c.ChannelID, fmt.Sprintf("o pana e toki kepeken open ni: `%s`", strings.Join(data, " ")), nil)

	gameStates[c.ChannelID] = sending
	dataStates[c.ChannelID] = sendData{
		letters: data,
		phrases: make(map[discord.UserID][]string),
	}

	go func() {
		time.Sleep(duration * time.Second)
		gameStates[c.ChannelID] = voting

		if len(dataStates[c.ChannelID].phrases) == 0 {
			bot.SendMessage(c.ChannelID, "jan ala li pana e toki ona! ike a...", nil)
			delete(gameStates, c.ChannelID)

			return
		}

		bot.SendMessage(c.ChannelID, "tenpo li pini a! o toki e wile sina tan ni:", nil)

		voteStates[c.ChannelID] = voteData{
			votes:    make(map[discord.UserID]int),
			hasVoted: make(map[discord.UserID]struct{}),
			keys:     make(map[int]discord.UserID),
		}

		s := strings.Builder{}
		i := 0

		for member, phrase := range dataStates[c.ChannelID].phrases {
			i++
			s.WriteString(fmt.Sprintf("%d - %s\n", i, strings.Join(phrase, " ")))
			voteStates[c.ChannelID].keys[i] = member
		}

		bot.SendMessage(c.ChannelID, s.String(), nil)

		go func() {
			time.Sleep(duration * time.Second)
			delete(gameStates, c.ChannelID)

			voteData := voteStates[c.ChannelID]
			phraseData := dataStates[c.ChannelID]

			delete(dataStates, c.ChannelID)
			delete(voteStates, c.ChannelID)

			winner := discord.UserID(0)
			votes := -1

			for user, voted := range voteData.votes {
				if voted > votes {
					winner = user
				}
			}

			if winner == 0 {
				bot.SendMessage(c.ChannelID, "jan ala li toki e wile sina. ike a...", nil)
				delete(gameStates, c.ChannelID)

				return
			}

			user, err := client.User.Query().Where(user.DiscordID(uint64(winner))).Only(ctx)
			if err != nil {
				if _, ok := err.(*ent.NotFoundError); ok {
					user = client.User.Create().
						SetDiscordID(uint64(winner)).
						SaveX(ctx)
				} else {
					panic(err)
				}
			}

			count := user.Update().AddWonGames(1).SaveX(ctx).WonGames

			bot.SendMessage(c.ChannelID, fmt.Sprintf("jan mute li wile toki pi <@%d>! ona li toki e ni: %s. ni li tenpo nanpa %d tawa ona.", winner, strings.Join(phraseData.phrases[winner], " "), count), nil)
		}()
	}()
}
