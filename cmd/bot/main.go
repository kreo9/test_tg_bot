package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kreo9/test_tg_bot/internal/service/product"
)

func main() {

	godotenv.Load("../../.env")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	productService := product.NewService()

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
		case "list":
			listCommand(bot, update.Message, productService)
		default:
			defaultBehavior(bot, update.Message)
		}

	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		"/help - help\n"+
			"/list - list products",
	)
	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message, productService *product.Service) {
	outputText := "All products: \n\n"
	products := productService.List()
	for _, p := range products {
		outputText += p.Title
		outputText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputText)
	bot.Send(msg)
}

func defaultBehavior(bot *tgbotapi.BotAPI, inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Я тупой бот, который только и умеет здороваться. Привет "+inputMsg.From.UserName+".")
	msg.ReplyToMessageID = inputMsg.MessageID
	bot.Send(msg)
}
