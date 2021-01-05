package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"iloMusiAko/ent"

	"github.com/diamondburned/arikawa/session"
	_ "github.com/mattn/go-sqlite3"
)

var bot *session.Session
var client *ent.Client

func main() {
	var err error

	client, err = ent.Open("sqlite3", "file:data.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	if err = client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

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
