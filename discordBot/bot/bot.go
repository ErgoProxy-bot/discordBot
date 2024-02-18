package bot

import (
	"fmt"
	"strings"

	"github.com/ErgoProxy-bot/discordBot"
	"github.com/ErgoProxy-bot/discordBot/config"
)

var BotId string
var goBot *discordBot.Session

func Start() {
	cfg, err := config.ReadConfig()
	if err != nil {
		fmt.Println("Failed reading config: ", err)
		return
	}

	goBot, err = discordBot.New("Bot " + cfg.Token)
	if err != nil {
		fmt.Println("Failed initializing Discord Session: ", err)
		return
	}

	u, err := goBot.User("@me")
	if err != nil {
		fmt.Println("Failed getting current user: ", err)
		return
	}

	BotId = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()
	if err != nil {
		fmt.Println("Failed opening connection to Discord: ", err)
		return
	}

	fmt.Println("Bot is now connected!")
}

func messageHandler(s *discordBot.Session, e *discordBot.MessageCreate) {
	if e.Author.ID == BotId {
		return
	}

	prefix := config.BotPrefix
	if strings.HasPrefix(e.Content, prefix) {
		args := strings.Fields(e.Content)[strings.Index(e.Content, prefix):]
		cmd := args[0][len(prefix):]
		arguments := args[1]

		switch cmd {
		case "ping":
			_, err := s.ChannelMessageSend(e.ChannelID, "Pong!")
			if err != nil {
				fmt.Println("Failed sending Pong response: ", err)
			}
		default:
			_, err := s.ChannelMessageSend(e.ChannelID, fmt.Sprintf("Unknown command %q.", cmd))
			if err != nil {
				fmt.Println("Failed sending unknown command response: ", err)
			}
			fmt.Println("Unhandled command arguments: ", arguments)
		}
	}
}
