package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type config struct {
	Bot struct {
		Token string `toml:"Token"`
	} `toml:"Bot"`
}

var botConfig config

func init() {
	println("Reading config...")
	_, err := toml.DecodeFile("config.toml", &botConfig)
	if err != nil {
		log.Fatalf("Failed to read config: %+v", err)
	}
}
