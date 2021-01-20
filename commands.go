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

func (s sendData) searchAlreadyHas(p []string) bool {
outer:
	for _, phrase := range s.phrases {
		if len(phrase) != len(p) {
			continue
		}
		for i := range phrase {
			if phrase[i] != p[i] {
				continue outer
			}
		}
		return true
	}
	return false
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

	if c.Author.ID == id {
		bot.SendMessage(c.ChannelID, fmt.Sprintf("ike a, <@%d> li pali ike!", c.Author.ID), nil)
		bot.DeleteMessage(c.ChannelID, c.ID)
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

func pona(title string, body string) *discord.Embed {
	return &discord.Embed{
		Color:       0x3dd425,
		Title:       title,
		Description: body,
	}
}

func ike(title string, body string) *discord.Embed {
	return &discord.Embed{
		Color:       0xe93d58,
		Title:       title,
		Description: body,
	}
}

func meso(title string, body string) *discord.Embed {
	return &discord.Embed{
		Color:       0xef973c,
		Title:       title,
		Description: body,
	}
}

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

	if strings.ContainsRune(c.Content, '⠀') {
		bot.SendMessage(c.ChannelID, fmt.Sprintf("<@%d> o, sitelen ni li ike MUTE a! o kepeken ALA e ona!", c.Author.ID), nil)
		bot.DeleteMessage(c.ChannelID, c.ID)
		return
	}

	if dataStates[c.ChannelID].searchAlreadyHas(fields) {
		bot.SendMessage(c.ChannelID, fmt.Sprintf("<@%d> o, jan ante li kepeken toki sina. o toki ante!", c.Author.ID), nil)
		bot.DeleteMessage(c.ChannelID, c.ID)
		return
	}

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

func leaderboard(c *gateway.MessageCreateEvent) {
	if c.WebhookID.IsValid() {
		return
	}

	if !strings.HasPrefix(c.Content, "ilo o leaderboard") {
		return
	}

	users := client.User.Query().
		Order(ent.Desc(user.FieldWonGames)).
		Limit(10).
		AllX(ctx)

	sb := []string{}

	for _, user := range users {
		sb = append(sb, fmt.Sprintf("<@%d> - tenpo %d", user.DiscordID, user.WonGames))
	}

	bot.SendMessage(
		c.ChannelID, "",
		pona(
			"jan mute li wile e toki pi jan ni:",
			strings.Join(sb, "\n"),
		),
	)
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

	bot.SendMessage(c.ChannelID, "", pona(fmt.Sprintf("o pana e toki kepeken open ni: `%s`", strings.Join(data, " ")), ""))

	gameStates[c.ChannelID] = sending
	dataStates[c.ChannelID] = sendData{
		letters: data,
		phrases: make(map[discord.UserID][]string),
	}

	go func() {
		go func() {
			time.Sleep(0.8 * duration * time.Second)
			bot.SendMessage(c.ChannelID, "tenpo li weka!", nil)
		}()

		time.Sleep(duration * time.Second)
		gameStates[c.ChannelID] = voting

		if len(dataStates[c.ChannelID].phrases) == 0 {
			bot.SendMessage(c.ChannelID, "", ike("jan ala li pana e toki ona! ike a...", ""))
			delete(gameStates, c.ChannelID)

			return
		}

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

		bot.SendMessage(c.ChannelID, "", pona("tenpo li pini a! o toki e wile sina tan ni:", s.String()))

		go func() {
			go func() {
				time.Sleep(0.8 * duration * time.Second)
				bot.SendMessage(c.ChannelID, "tenpo li weka!", nil)
			}()

			time.Sleep(duration * time.Second)
			defer delete(gameStates, c.ChannelID)

			voteData := voteStates[c.ChannelID]
			phraseData := dataStates[c.ChannelID]

			defer delete(dataStates, c.ChannelID)
			defer delete(voteStates, c.ChannelID)

			winner := discord.UserID(0)
			couldWinner := discord.UserID(0)
			votes := -1
			couldVotes := -1

			for user, voted := range voteData.votes {
				_, ok := voteStates[c.ChannelID].hasVoted[user]

				if voted > couldVotes {
					couldWinner = user
					couldVotes = voted
				}
				if voted > votes && ok {
					winner = user
					votes = voted
				}
			}

			if winner == 0 {
				if couldVotes <= votes {
					bot.SendMessage(c.ChannelID, "", ike("jan ala li toki e wile ona. ike a...", ""))
				} else {
					bot.SendMessage(c.ChannelID, "", meso("", fmt.Sprintf("tenpo ante la <@%d> li ken pona, taso ona li toki ala e wile ona. jan li toki ala e wile tawa jan ante.", couldWinner)))
				}
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

			if couldVotes <= votes {
				bot.SendMessage(c.ChannelID, "", pona("", fmt.Sprintf("jan mute li wile toki pi <@%d>!\nona li toki e ni: %s.\nni li tenpo nanpa %d tawa ona.", winner, strings.Join(phraseData.phrases[winner], " "), count)))
			} else {
				bot.SendMessage(c.ChannelID, "", pona("", fmt.Sprintf("jan mute li wile toki pi <@%d>!\nona li toki e ni: %s.\nni li tenpo nanpa %d tawa ona.\ntenpo ante la <@%d> li ken pona, taso ona li toki ala e wile ona.", winner, strings.Join(phraseData.phrases[winner], " "), count, couldWinner)))
			}
		}()
	}()
}
