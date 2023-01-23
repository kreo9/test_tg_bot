package commands

import (
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
