package main

import (
	"github.com/ErgoProxy-bot/discordBot/bot"
)

func main() {
	bot.Start()

	<-make(chan struct{})
}
