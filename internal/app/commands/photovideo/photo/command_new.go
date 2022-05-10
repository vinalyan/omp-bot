package photo

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/photovideo"
)

func (c *PhotoCommander) New(inputMsg *tgbotapi.Message) {

	args := inputMsg.CommandArguments()
	var msg tgbotapi.MessageConfig
	_, err := c.photoService.Create(*photovideo.NewPhoto(0, args))
	if err != nil {
		msg = tgbotapi.NewMessage(inputMsg.Chat.ID, "Обшибка добавления товара")
		log.Printf("PhotoCommander.List: error sending reply message to chat - %v", err)
	} else {
		msg = tgbotapi.NewMessage(inputMsg.Chat.ID, "Товар успешно добавлен")
	}
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PhotoCommander.Get: error sending reply message to chat - %v", err)
	}

}
