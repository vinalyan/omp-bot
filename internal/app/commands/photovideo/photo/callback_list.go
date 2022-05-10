package photo

import (
	"encoding/json"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset int `json:"offset"`
}

func (c *PhotoCommander) CallbackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	const limitOnPage = 4

	parsedData := CallbackListData{}
	err := json.Unmarshal([]byte(callbackPath.CallbackData), &parsedData)
	if err != nil {
		log.Printf("PhotoCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	fmt.Println("OFFSET: ", parsedData.Offset)

	products, err := c.photoService.List(uint64(parsedData.Offset), uint64(limitOnPage))
	if err != nil {
		log.Println("PhotoCommander.List.CallbackList: List is empty")
	}

	viewSize := limitOnPage
	if len(products) < limitOnPage {
		viewSize = len(products)
	}

	outputMsgText := fmt.Sprintf("List of the products (from %v to %v)\n\n",
		parsedData.Offset+1,
		parsedData.Offset+viewSize,
	)

	for _, p := range products {
		outputMsgText += p.String()
		outputMsgText += "\n"
	}

	//msg := tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	var msg tgbotapi.MessageConfig
	if len(products) == 0 {
		msg = tgbotapi.NewMessage(callback.Message.Chat.ID, "No more products")
		_, err = c.bot.Send(msg)
		if err != nil {
			log.Printf("PhotoCommander.CallbackList: error sending reply message to chat - %v", err)
		}
		return
	} else {
		msg = tgbotapi.NewMessage(callback.Message.Chat.ID, outputMsgText)
	}

	serializedData, _ := json.Marshal(CallbackListData{
		Offset: parsedData.Offset + limitOnPage,
	})

	сallbackPath := path.CallbackPath{
		Domain:       "photovideo",
		Subdomain:    "photo",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Next page", сallbackPath.String()),
		),
	)

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DummyPhotoCommander.CallbackList: error sending reply message to chat - %v", err)
	}
}
