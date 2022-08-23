// Package utils contains utility functions needed throughout the app
package utils

import (
	"encoding/json"
	"os"

	logger "github.com/withmandala/go-log"
)

// Prefix is the bot command character prefix
const Prefix string = "?"

type configStruct struct {
	Token string `json:"Token"`
}

var (
	conf *configStruct
	log  = NewLogger()
)

// NewLogger returns a new instance of a logger
func NewLogger() *logger.Logger {
	return logger.New(os.Stdout).WithColor()
}

// ReadConfig reads the config file for the bot Token
func ReadConfig(filepath string) (string, error) {
	log.Info("Looking for discord token...")
	// Read file content
	content, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	// Unmarshal content into config struct
	err = json.Unmarshal(content, &conf)
	if err != nil {
		return "", err
	}

	log.Info("Discord token received successfully!")
	return conf.Token, nil
}
