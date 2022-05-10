package photo

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	photoservice "github.com/ozonmp/omp-bot/internal/service/photovideo/photo"
)

type PhotoCommander struct {
	bot          *tgbotapi.BotAPI
	photoService photoservice.PhotoService
}

func NewDummyPhotoCommander(bot *tgbotapi.BotAPI, photoService photoservice.PhotoService) *PhotoCommander {

	return &PhotoCommander{
		bot:          bot,
		photoService: photoService,
	}
}

func (c *PhotoCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallbackList(callback, callbackPath)
	default:
		log.Printf("DemoSubdomainCommander.HandleCallback: unknown callback name: %s", callbackPath.CallbackName)
	}
}

func (c *PhotoCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(msg)
	case "get":
		c.Get(msg)
	case "list":
		c.List(msg)
	case "delete":
		c.Delete(msg)
	case "new":
		c.New(msg)
	case "edit":
		c.Edit(msg)
	default:
		c.Default(msg)
	}
}
