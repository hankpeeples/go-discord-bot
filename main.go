// Package main is the 'driver' of the discord bot
package main

import (
	_ "fmt"
	"os"

	"github.com/hankpeeples/go-discord-bot/utils"
)

func main() {
	// set bot command prefix
	const PREFIX = ">"

	// Make sure config.json file was given as an argument
	if len(os.Args) != 2 {
		utils.Logger.Fatalf("Usage: ./{binary} config.json")
	}

	// Read config file for bot Token
	_, err := utils.ReadConfig(os.Args[1])
	if err != nil {
		utils.Logger.Fatalf("ReadConfig(): %v\n", err)
	}
}
