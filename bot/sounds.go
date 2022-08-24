package bot

import (
	"encoding/binary"
	"io"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

var airhornBuffer = make([][]byte, 0)
var xGamesModeBuffer = make([][]byte, 0)

// LoadSounds loads in the .dca sound files
func LoadSounds() {
	log.Info("Loading sound assets...")

	airhorn, err := os.Open("assets/airhorn.dca")
	if err != nil {
		log.Errorf("Unable to open airhorn.dca: %s", err)
	}

	xgames, err := os.Open("assets/x-games-mode.dca")
	if err != nil {
		log.Errorf("Unable to open x-games-mode.dca: %s", err)
	}

	readSoundFile("airhorn", airhorn)
	readSoundFile("x-games-mode", xgames)
}

func readSoundFile(name string, file *os.File) error {
	buffer := make([][]byte, 0)
	var opuslen int16

	for {
		// Read opus frame length from dca file
		err := binary.Read(file, binary.LittleEndian, &opuslen)

		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err := file.Close()
			if err != nil {
				log.Errorf("Unable to close %s.dca file: %s", name, err)
				return err
			}
			break
		}

		if err != nil {
			log.Errorf("Unable to read %s.dca file: %s", name, err)
			return err
		}

		// Read encoded pcm from dca file
		InBuf := make([]byte, opuslen)
		err = binary.Read(file, binary.LittleEndian, &InBuf)
		if err != nil {
			log.Errorf("Unable to read %s.dca file: %s", name, err)
			return err
		}

		// Append encoded pcm data to buffer
		buffer = append(buffer, InBuf)
	}

	if name == "airhorn" {
		airhornBuffer = buffer
	} else if name == "x-games-mode" {
		xGamesModeBuffer = buffer
	}

	return nil
}

// PlaySound plays the airhorn sound in the callers voice channel
func PlaySound(s *discordgo.Session, guildID, channelID string, sound string) error {
	buffer := make([][]byte, 0)

	if sound == "airhorn" {
		buffer = airhornBuffer
	} else if sound == "x-games-mode" {
		buffer = xGamesModeBuffer
	}

	// Join voice channel
	vc, err := s.ChannelVoiceJoin(guildID, channelID, false, true)
	if err != nil {
		log.Error("Unable to join voice channel: ", err)
		return err
	}

	// Sleep before playing sound
	time.Sleep(250 * time.Millisecond)

	// start speaking
	vc.Speaking(true)

	// Send buffered data
	for _, buff := range buffer {
		vc.OpusSend <- buff
	}

	// stop speaking
	vc.Speaking(false)

	// Disconnect from voice channel
	time.Sleep(250 * time.Millisecond)
	vc.Disconnect()

	return nil
}
