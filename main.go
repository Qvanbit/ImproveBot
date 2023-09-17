package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var gBot *tgbotapi.BotAPI
var gToken string
var gChatID int64

func isStartMessage(update *tgbotapi.Update) bool {
	return update.Message != nil && update.Message.Text == "/start"
}

func init() {
	_ = os.Setenv(TOKEN_NAME_ABOUT_BOT, TOKEN)

	gToken = os.Getenv(TOKEN_NAME_ABOUT_BOT)
	var err error
	if gBot, err = tgbotapi.NewBotAPI(gToken); err != nil {
		log.Panic(err)
	}

	gBot.Debug = true
}

func main() {

	log.Printf("Authorized on account %s", gBot.Self.UserName)

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UPDATE_CONFIG_TIMEOUT

	for update := range gBot.GetUpdatesChan(updateConfig) {
		if isStartMessage(&update) {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			gBot.Send(msg)
		}
	}
}
