package photo

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

func (c *PhotoCommander) List(inputMsg *tgbotapi.Message) {
	//outputMsgText := "Here all the products: \n\n"

	const firstIndx = 0
	const limitOnPage = 4

	products, err := c.photoService.List(uint64(firstIndx), uint64(limitOnPage))
	if err != nil {
		log.Printf("PhotoCommander.List: error sending List - %v", err)
	}

	viewSize := limitOnPage
	if len(products) < limitOnPage {
		viewSize = len(products)
	}
	outputMsgText := fmt.Sprintf("List of the products (from %v to %v)\n\n", 1, viewSize)
	for _, p := range products {
		outputMsgText += p.String()
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
		log.Printf("PhotoCommander.List: error sending reply message to chat - %v", err)
	}
}
