package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/hankpeeples/go-discord-bot/utils"
)

// HelpCommand is run from the 'help' command
func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Commands",
		Description: "Use '" + utils.Prefix + "' before each command.\nClick the title to see the code that created me.",
		URL:         "https://github.com/hankpeeples/go-discord-bot",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "help",
				Value: "Shows this message!",
			},
			{
				Name:  "latency",
				Value: "Shows the bots current latency to the server.",
			},
			{
				Name:  "airhorn",
				Value: "Plays airhorn sound. (Must be in a voice channel)",
			},
			{
				Name:  "x-games-mode",
				Value: "Plays x games mode sound. (Must be in a voice channel)",
			},
			{
				Name:  "goofy",
				Value: "Plays goofy trials sound. (Must be in a voice channel)",
			},
		},
		Color:  blue,
		Footer: footer,
	})
	if err != nil {
		utils.HandleEmbedFailure(s, m, err)
	}
}

// LatencyCommand is run from the 'latency' command
func LatencyCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Latency to server",
		Description: fmt.Sprint(s.HeartbeatLatency()),
		Color:       green,
		Footer:      footer,
	})
	if err != nil {
		utils.HandleEmbedFailure(s, m, err)
	}
}

// AirhornCommand is run from the 'airhorn' command
func AirhornCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := initializeSound(s, m, "airhorn")
	if err == nil {
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: "Playing airhorn sound",
			Color:       blue,
		})
		if err != nil {
			utils.HandleEmbedFailure(s, m, err)
		}
	}
}

// XgamesModeCommand is run from the 'x-games-mode' command
func XgamesModeCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := initializeSound(s, m, "x-games-mode")
	if err == nil {
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: "Playing x games mode sound",
			Color:       blue,
		})
		if err != nil {
			utils.HandleEmbedFailure(s, m, err)
		}
	}
}

// GoofyCommand is run from the 'goofy` command`
func GoofyCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	err := initializeSound(s, m, "goofy")
	if err == nil {
		_, err = s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: "Playing goofy trials sound",
			Color:       blue,
		})
		if err != nil {
			utils.HandleEmbedFailure(s, m, err)
		}
	}
}
