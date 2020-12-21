package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/diamondburned/arikawa/session"
)

var bot *session.Session

func main() {
	var err error

	bot, err = session.New("Bot " + botConfig.Bot.Token)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	bot.AddHandler(startGame)
	bot.AddHandler(dataPhase)
	bot.AddHandler(votePhase)

	if err := bot.Open(); err != nil {
		log.Fatalf("%+v", err)
	}
	defer bot.Close()

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc
}
