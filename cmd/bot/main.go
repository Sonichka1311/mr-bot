package main

import (
	"context"
	"log"
	"mr-bot/pkg/basics"
	"os"
)

func main() {
	bot := basics.Bot{
		BotToken: os.Getenv("bot_token"),
	}
	err := bot.StartBot(context.Background())
	if err != nil {
		log.Fatalf(err.Error())
	}
}
