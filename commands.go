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
    "unicode"

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

const dur = 70

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
		bot.SendMessage(c.ChannelID, fmt.Sprintf("ike a · <@%d> li pali ike ·", c.Author.ID), nil)
		bot.DeleteMessage(c.ChannelID, c.ID)
		return
	}

	count, _ := voteStates[c.ChannelID].votes[id]

	voteStates[c.ChannelID].votes[id] = count + 1
	voteStates[c.ChannelID].hasVoted[c.Author.ID] = struct{}{}

	bot.DeleteMessage(c.ChannelID, c.ID)
	bot.SendMessage(c.ChannelID, fmt.Sprintf("<@%d> li pana e wile ona ·", c.Author.ID), nil)
}

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

// strips punctuation from string
// Don’t implement Unicode folding here.
// Only fold in the storing function, where it will be checked against others’ submissions.
func strip (s string) string {
    var ns string // 
    var weka bool = true

    for _, mu := range s { // mu ali la
        if unicode.IsLetter(mu) || unicode.IsNumber(mu) {
            if weka {
                ns += ' ' // ken la " " li pona
                weka = false
            }
            ns += string(mu)
        } else {
            weka = true
        }
    }
    
    // will make «nimi ‹kokosila› li ike · » into « nimi kokosila li ike»
    // this will now be put into checksOut
    // if it checksOut, then it will be Unicode-folded and stored
    
    return ns[1:] // with first space removed (is always a space)
}

type validationState int

const (
	ala validationState = iota
	ken
	lon
)

type statedWord struct {
    word  string
    state validationState
}

var letters []string   // What do these variables do? Have I written them?
var sentence []string  // They seem to be unused. But why does Go not raise an error for this? – jan Kasape

func statedWordsFromSentence(s []string) []statedWord {
    sw := make([]statedWord, len(s))
    
    for i := 0; i < len(s); i++ {
        sw[i] = statedWord { word: s[i], state: ala, }
    }
    
    return sw
}

// s = every word in the input accompanied by a ala-ken-lon state
// l = array of original letters, e.g. {"a", "k", "o"}
func checksOut(s []statedWord, l []string) bool {

    var pl int = 0 // letter pointer
    var ps int = 0 // sentence word pointer
    
    // returns true iff match between letters and sentence exists
    for {
        var es bool = ps >= len(s) // exceeded sentence pointer
        var el bool = pl >= len(l) // exceeded letter pointer
    
        /* /// pointers and flags check
        fmt.Println(pl, ps)
        for i := 0; i < len(s); i++ {
            fmt.Print(s[i].state, " ")
        }
        fmt.Println()
        */
        
        if el && es {
            return true // exceeded both? then there exists a mapping
        }
        
        if !es {
            // current word’s first letter == current letter (case-insensitive)
            var sama bool
            
            // if-block in case there are particles at the end of the input
            if el { 
                sama = false // if exceeded, evidently can’t match
            } else {
                sama = strings.EqualFold(string(s[ps].word[0]), l[pl])
            }
            
            var pana bool = isPart(s[ps].word) // is free word

            if /****/ !sama &&  pana {
                s[ps].state = ala
                ps++ // particle doesn’t match letter: skip particle
                continue
            
            } else if  sama && !pana {
                s[ps].state = lon
                ps++ // matches a non-particle: has to count as letter
                pl++
                continue
            
            } else if  sama &&  pana {
                // could be that the particle should be counted
                // or not, so we try both; first skip the particle
                // and then, perhaps later, count the particle
                s[ps].state = ken
                ps++
                continue
            }
        }
        
        // if here: mismatch
        // go back to the rightmost ‹ken› marker and mark as ‹lon›
        for {
            if pl < 0 {
                return false
            
            } else if ps >= len(s) ||
                      s[ps].state == lon ||
                      s[ps].state == ala { // undo last action …
                ps--
                if ps < 0 {
                    return false // (quick check)
                }
                
                if s[ps].state == lon { // … so, if from non-particle phase
                    pl--                // then also decrease letter pointer
                }
            } else if s[ps].state == ken {
                s[ps].state = lon
                ps++
                pl++
                break
            }
            /// fmt.Println("·", pl, ps)
        }
    }
}
/// 1
// If both counters exceeded, then the word will be accepted.
// If only the sentence pointer exceeded,
// then there doesn’t exist a matching.
// Because, since we accept particles by default,
// the accepted sentence will always be as least as long
// as the number of letters. Thus, if you’re in the middle
// of a sentence, and you’ve run out of letters already,
// a full match will never occur.


func dataPhase(c *gateway.MessageCreateEvent) {
	if c.WebhookID.IsValid() {
		return
	}

	if state, ok := gameStates[c.ChannelID]; ok && state == sending {
	} else {
		return
	}

	fields := strings.Fields(c.Content)

	if strings.ContainsRune(c.Content, '⠀') {
        bot.SendMessage(c.ChannelID, fmt.Sprintf("<@%d> o : sitelen ni li ike suli a · o kepeken ala ona ·", c.Author.ID), nil)
		bot.DeleteMessage(c.ChannelID, c.ID)
		return
	}

	if dataStates[c.ChannelID].searchAlreadyHas(fields) {
		bot.SendMessage(c.ChannelID, fmt.Sprintf("<@%d> o, jan ante li kepeken toki sina. o toki ante!", c.Author.ID), nil)
		bot.DeleteMessage(c.ChannelID, c.ID)
		return
	}

	s := statedWordsFromSentence(strip(c.Content)) // should work if c.Content is a string
	l := dataStates[c.ChannelID].letters
	if !checksOut(s, l) {
		return
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

	if !(strings.Contains(c.Content, "ilo o open e musi Ako") ||
		strings.Contains(c.Content, "ilo o Ako")) {
		return
	}

	var data []string

	duration := float64(dur)
	if strings.Contains(c.Content, "tenpo mute") {
		duration = dur * 1.5
		data = randomLetters(2)
	} else if strings.Contains(c.Content, "tenpo lili") {
		duration = dur * 0.5
		data = randomLetters(-1)
	} else if strings.Contains(c.Content, "tenpo pi lili mute") {
		duration = dur * 0.2
		data = randomLetters(-2)
	} else if strings.Contains(c.Content, "tenpo ale") {
		duration = 300
		data = randomLetters(12)
	} else {
		data = randomLetters(0)
	}

	if _, ok := gameStates[c.ChannelID]; ok {
		bot.SendMessage(c.ChannelID, "musi li lon. o musi kepeken ona!", nil)
		return
	}

	bot.SendMessage(c.ChannelID, "", pona(fmt.Sprintf("o pana e toki kepeken open ni: `%s`", strings.Join(data, " ")), ""))

	gameStates[c.ChannelID] = sending
	dataStates[c.ChannelID] = sendData{
		letters: data,
		phrases: make(map[discord.UserID][]string),
	}

	go func() {
		go func() {
			time.Sleep(time.Duration(0.8*duration) * time.Second)
			bot.SendMessage(c.ChannelID, "tenpo li weka!", nil)
		}()

		time.Sleep(time.Duration(duration) * time.Second)
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
				time.Sleep(time.Duration(0.8*duration) * time.Second)
				bot.SendMessage(c.ChannelID, "tenpo li weka!", nil)
			}()

			time.Sleep(time.Duration(duration) * time.Second)
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
