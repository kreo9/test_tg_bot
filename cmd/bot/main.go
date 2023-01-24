package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"github.com/kreo9/test_tg_bot/internal/app/commands"
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

	commander := commands.NewCommandRouter(bot, productService)
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		commander.HandleUpdate(update)
	}
}
