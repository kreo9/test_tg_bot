package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *CommandRouter) Default(inputMsg *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Я тупой бот, который только и умеет здороваться. Привет "+inputMsg.From.UserName+".")
	msg.ReplyToMessageID = inputMsg.MessageID
	c.bot.Send(msg)
}
