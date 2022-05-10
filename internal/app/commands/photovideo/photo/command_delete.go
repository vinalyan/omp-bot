package photo

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *PhotoCommander) Delete(inputMsg *tgbotapi.Message) {

	args := inputMsg.CommandArguments()
	var msg tgbotapi.MessageConfig
	idx, err := strconv.Atoi(args)
	//TODO добавить описание удаляемого объекта
	if err != nil || idx <= 0 {
		log.Println("wrong args", args)
		msg = tgbotapi.NewMessage(inputMsg.Chat.ID, "wrong args")
	} else {
		_, err := c.photoService.Remove(uint64(idx))
		if err != nil {
			log.Printf("PhotoCommander.Get: error get item - %v", err)
			msg = tgbotapi.NewMessage(inputMsg.Chat.ID, "Значение не найдено")
		} else {
			msg = tgbotapi.NewMessage(inputMsg.Chat.ID, "Объект удален")
		}
	}

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("PhotoCommander.Get: error sending reply message to chat - %v", err)
	}
}
