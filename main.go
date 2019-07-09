package main

import (
	"log"
	"os"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	chat := os.Getenv("CHAT")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	log.Println("Bot authorized")
	log.Printf("Start conversation: https://t.me/%s", bot.Self.UserName)

	msg := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChannelUsername:  chat,
			ReplyToMessageID: 0,
		},
		Text:                  strings.Join(os.Args[1:], "\n"),
		DisableWebPagePreview: false,
	}

	_, err = bot.Send(msg)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Message sent")
}
