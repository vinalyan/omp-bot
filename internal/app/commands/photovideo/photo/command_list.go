package photo

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *PhotoCommander) List(inputMsg *tgbotapi.Message) {
	outputMsgText := "Here all the products: \n\n"

	firstIndx := 0
	//entitesCount := c.photoService
	limitOnPage := 4

	products, err := c.photoService.List(uint64(firstIndx), uint64(limitOnPage))
	if err != nil {
		log.Printf("DummyPhotoCommander.List: error sending List - %v", err)
	}
	for _, p := range products {
		outputMsgText += p.Name
		outputMsgText += "\n"
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID, outputMsgText)

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: limitOnPage,
	})

	callbackPath := path.CallbackPath{
		Domain:       "photovideo",
		Subdomain:    "photo",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", callbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyPhotoCommander.List: error sending reply message to chat - %v", err)
	}
}
