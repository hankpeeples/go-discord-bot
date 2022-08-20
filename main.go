// Package main is the 'driver' of the discord bot
package main

import (
	_ "fmt"
	"os"

	"github.com/hankpeeples/go-discord-bot/config"
	"github.com/withmandala/go-log"
)

func main() {
	// initialize logger
	var logger *log.Logger = log.New(os.Stdout).WithColor().WithoutTimestamp()
	// set bot command prefix
	const PREFIX = ">"

	// Make sure config.json file was given as an argument
	if len(os.Args) != 2 {
		logger.Fatalf("Usage: ./{binary} config.json")
	}

	// Read config file for bot Token
	Token, err := config.ReadConfig(os.Args[1])
	if err != nil {
		logger.Fatalf("ReadConfig(): %v\n", err)
	}

	logger.Info(Token)
}
