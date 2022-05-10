package photo

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	service "github.com/ozonmp/omp-bot/internal/service/photovideo/photo"
)

type PhotoCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)

	New(inputMsg *tgbotapi.Message)  // return error not implemented
	Edit(inputMsg *tgbotapi.Message) // return error not implemented
}

type Commander struct {
}

func NewPhotoCommander(bot *tgbotapi.BotAPI, service service.PhotoService) Commander {
	// ...
	return Commander{}
}
