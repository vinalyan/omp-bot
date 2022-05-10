package photovideo

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	photocommander "github.com/ozonmp/omp-bot/internal/app/commands/photovideo/photo"
	"github.com/ozonmp/omp-bot/internal/app/path"
	"github.com/ozonmp/omp-bot/internal/service/photovideo/photo"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}

type PhotovideoCommander struct {
	bot            *tgbotapi.BotAPI
	photoCommander photocommander.PhotoCommander
}

func NewPhotovideoCommander(bot *tgbotapi.BotAPI) *PhotovideoCommander {
	return &PhotovideoCommander{
		bot:            bot,
		photoCommander: photocommander.NewDummyPhotoCommander(bot, photo.NewDummyPhotoService()),
	}
}

func (c *PhotovideoCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "photo":
		c.photoCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("PhotovideoCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *PhotovideoCommander) HandleCommand(msg *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "photo":
		c.photoCommander.HandleCommand(msg, commandPath)
	default:
		log.Printf("PhotovideoCommander.HandleCommand: unknown subdomain - %s", commandPath.Subdomain)
	}
}
