// Package utils contains utility functions needed throughout the app
package utils

import (
	"encoding/json"
	"os"

	"github.com/withmandala/go-log"
)

// Logger initialization
var Logger *log.Logger = log.New(os.Stdout).WithColor()

// Prefix is the bot command character prefix
const Prefix string = ">"

type configStruct struct {
	Token string `json:"Token"`
}

var conf *configStruct

// ReadConfig reads the config file for the bot Token
func ReadConfig(filepath string) (string, error) {
	Logger.Info("Looking for discord token...")
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

	Logger.Info("Discord token received successfully!")
	return conf.Token, nil
}
