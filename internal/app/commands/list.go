package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *CommandRouter) List(inputMsg *tgbotapi.Message) {
	outputText := "Все продукты: \n\n"
	products := c.productService.List()
	for _, p := range products {
		outputText += p.Title
		outputText += "\n"
	}
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputText)
	c.bot.Send(msg)
}
