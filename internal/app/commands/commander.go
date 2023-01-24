package commands

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kreo9/test_tg_bot/internal/service/product"
)

type CommandRouter struct {
	bot            *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommandRouter(bot *tgbotapi.BotAPI, productService *product.Service) *CommandRouter {
	return &CommandRouter{
		bot:            bot,
		productService: productService,
	}
}

func (c *CommandRouter) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic %v", panicValue)
		}
	}()

	if update.Message == nil {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	switch update.Message.Command() {
	case "help":
		c.Help(update.Message)
	case "list":
		c.List(update.Message)
	case "get":
		c.Get(update.Message)
	default:
		c.Default(update.Message)
	}
}
