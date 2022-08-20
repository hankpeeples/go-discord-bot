// Package bot will contain bot specific functions
package bot

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/hankpeeples/go-discord-bot/utils"
)

// Start will begin a new discord bot session
func Start(token string) {
	utils.Logger.Info("Attempting to start bot session...")
	// Create discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		utils.Logger.Fatalf("Error creating discord session: %v", err)
	}

	// Register messageCreate func as callback for message events
	dg.AddHandler(messageCreate)

	// Need information about guilds (which includes channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open websocket and begin listening
	err = dg.Open()
	if err != nil {
		utils.Logger.Fatalf("Error opening websocket: %v", err)
	}

	utils.Logger.Info("Session open and listening ✅")
	// Wait here for ctrl-c or other termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	utils.Logger.Warn("Session terminated!")
	// close discordgo session after kill signal is received
	dg.Close()
}

// messageCreate will be called every time a new message is sent
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	utils.Logger.Infof("[%s]: %s", m.Author, m.Content)

	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if len(m.Content) >= 1 {
		// Only look for commands that begin with defined prefix character
		if m.Content[0:1] == utils.Prefix {
			// Set command to the 'command' after the prefix character
			command := m.Content[1:]
			// If the message is "ping" reply with "Pong!"
			if command == "ping" {
				s.ChannelMessageSend(m.ChannelID, "Pong!")
			}

			// If the message is "pong" reply with "Ping!"
			if command == "pong" {
				s.ChannelMessageSend(m.ChannelID, "Ping!")
			}
		}
	}
}
