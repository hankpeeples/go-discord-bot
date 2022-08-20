// Package config reads the bot configuration from the (hidden) config.json file.
package config

import (
	"encoding/json"
	"os"
)

type configStruct struct {
	Token string `json:"Token"`
}

var conf *configStruct

// ReadConfig reads the config file for the bot Token
func ReadConfig(filepath string) (string, error) {
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

	return conf.Token, nil
}
