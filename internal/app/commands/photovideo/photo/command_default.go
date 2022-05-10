package photo

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PhotoCommander) Default(inputMsg *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMsg.From.UserName, inputMsg.Text)

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, "Неизвестная команда")

	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("PhotoCommander.Default: error sending reply message to chat - %v", err)
	}
}
