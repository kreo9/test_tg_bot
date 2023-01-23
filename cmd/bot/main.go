package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load("../../.env")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	//bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		// Получили сообщение
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		switch update.Message.Command() {
		case "help":
			helpCommand(bot, update.Message)
		default:
			defaultBehavior(bot, update.Message)
		}

	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "/help - help")
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Я тупой бот, который только и умеет здороваться. Привет "+inputMsg.From.UserName+".")
	msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}
