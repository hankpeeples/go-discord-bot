package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/hankpeeples/go-discord-bot/utils"
)

// HelpCommand runs the 'help' command
func HelpCommand(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Commands",
		Description: "Use '" + utils.Prefix + "' before each command. These commands ARE case sensitive. \nClick the title to see the code that created me.",
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
