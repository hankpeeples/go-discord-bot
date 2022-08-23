// Package main is the 'driver' of the discord bot
package main

import (
	_ "fmt"
	"os"

	"github.com/hankpeeples/go-discord-bot/bot"
	"github.com/hankpeeples/go-discord-bot/utils"
)

var log = utils.NewLogger()

func main() {
	// Make sure config.json file was given as an argument
	if len(os.Args) != 2 {
		log.Fatalf("Usage: ./{binary} config.json")
	}

	// Read config file for bot Token
	token, err := utils.ReadConfig(os.Args[1])
	if err != nil {
		log.Fatalf("ReadConfig(): %v\n", err)
	}

	// Start Bot
	bot.Start(token)
}
