package bot

import (
	"encoding/binary"
	"io"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var buffer = make([][]byte, 0)

// LoadAirhorn loads in the airhorn.dca sound file
func LoadAirhorn() error {
	log.Info("Loading airhorn...")

	file, err := os.Open("assets/airhorn.dca")
	if err != nil {
		log.Errorf("Unable to open airhorn.dca file: %s", err)
		return err
	}

	var opuslen int16

	for {
		// Read opus frame length from dca file
		err = binary.Read(file, binary.LittleEndian, &opuslen)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				log.Errorf("Unable to close airhorn.dca file: %s", err)
				return err
			}
			return nil
		}

		if err != nil {
			log.Errorf("Unable to read airhorn.dca file: %s", err)
			return err
		}

		// Read encoded pcm from dca file
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)
		if err != nil {
			log.Errorf("Unable to read airhorn.dca file: %s", err)
			return err
		}

		// Append encoded pcm data to buffer
		buffer = append(buffer, InBuf)
	}
}

// PlaySound plays the airhorn sound in the callers voice channel
func PlaySound(s *discordgo.Session, guildID, channelID string) error {
	// Join voice channel
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		log.Error("Unable to join voice channel: ", err)
		return err
	}

	// Sleep before playing sound
	time.Sleep(200 * time.Millisecond)

	// start speaking
	vc.Speaking(true)

	// Send buffered data
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// stop speaking
	vc.Speaking(false)

	// Disconnect from voice channel
	time.Sleep(200 * time.Millisecond)
	vc.Disconnect()

	return nil
}
