// Package bot will contain bot specific functions
package bot

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/hankpeeples/go-discord-bot/utils"
)

var (
	log    = utils.NewLogger()
	red    = 0xf54248
	blue   = 0x42b9f5
	green  = 0x28de4f
	footer = &discordgo.MessageEmbedFooter{
		Text: "Last bot reboot: " + time.Now().Format("Mon, 02 Jan 2006 15:04:05"),
	}
)

// Start will begin a new discord bot session
func Start(token string) {
	log.Info("Attempting to start bot session...")
	// Create discord session
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("Error creating discord session: %v", err)
	}

	// Load sounds
	LoadSounds()

	// Register ready func as callback for ready events
	dg.AddHandler(ready)
	// Register messageCreate func as callback for message events
	dg.AddHandler(messageCreate)

	// Need information about guilds (which includes channels),
	// messages and voice states.
	dg.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages | discordgo.IntentsGuildVoiceStates

	// Open websocket and begin listening
	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening websocket: %v", err)
	}

	log.Info("Session open and listening ✅")
	// Wait here for ctrl-c or other termination signal
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	log.Warn("Session terminated!")
	// close discordgo session after kill signal is received
	dg.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	// Set status message
	err := s.UpdateListeningStatus(utils.Prefix + "help")
	if err != nil {
		log.Error("Status message was NOT updated...")
		return
	}
	log.Info("Bot status message updated successfully")
}

// messageCreate will be called every time a new message is sent
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	if len(m.Content) >= 1 {
		// Only look for commands that begin with defined prefix character
		if strings.HasPrefix(m.Content, utils.Prefix) && (m.ChannelID == "706588663747969107" || m.ChannelID == "706431216336896012") {
			// Log received commands
			log.Infof("[%s]: %s", m.Author, m.Content)

			// Split string by whitespace (after prefix), for commands
			// that can accept arguments
			command := strings.Fields(m.Content[1:])
			// Set args to the strings after the command name
			var args []string
			for i, arg := range command {
				// skip command name
				if i == 0 {
					continue
				}
				args = append(args, arg)
			}

			switch command[0] {
			case "help":
				HelpCommand(s, m)
				break

			case "latency":
				LatencyCommand(s, m)
				break

			case "airhorn":
				AirhornCommand(s, m)
				break

			case "x-games-mode":
				XgamesModeCommand(s, m)
				break

			case "goofy":
				GoofyCommand(s, m)
				break

			default:
				_, err := s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
					Description: fmt.Sprintf("`%s` is not a command...", command[0]),
					Color:       red,
				})
				if err != nil {
					utils.HandleEmbedFailure(s, m, err)
				}
				break
			}
		} else {
			if strings.HasPrefix(m.Content, utils.Prefix) {
				_, err := s.ChannelMessageSendEmbedReply(m.ChannelID, &discordgo.MessageEmbed{
					Description: "Please use bot commands in the `bot-commands` text channel.",
					Color:       red,
				}, &discordgo.MessageReference{
					ChannelID: m.ChannelID,
					MessageID: m.ID,
				})
				if err != nil {
					utils.HandleEmbedFailure(s, m, err)
				}
			}
		}
	}
}
