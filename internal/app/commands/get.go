package commands

import (
	"fmt"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *CommandRouter) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()
	n, err := strconv.Atoi(args)
	if err != nil {
		log.Println("unknown arg", args)
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
			"Не получилось распознать текст после команды.")
		c.bot.Send(msg)
		return
	}

	product, err := c.productService.Get(n)
	if err != nil {
		log.Printf("fail to get product with idx %v: %v", n, err)
		msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
			"Такого продукта не существует.")
		c.bot.Send(msg)
		// тест
		panic("index out of range")
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		fmt.Sprintf("%v", product.Title))
	c.bot.Send(msg)
}
