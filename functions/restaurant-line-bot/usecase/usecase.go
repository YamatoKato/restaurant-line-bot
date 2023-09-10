package usecase

import (
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/repository"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type IUsecase interface {
	GetRestaurantInfos(apiParams *model.APIParams) (*linebot.TemplateMessage, error)
	SetAreaMenu() (*linebot.TemplateMessage, error)
	SetGenreMenu(apiParams *model.APIParams) (*linebot.TemplateMessage, error)
}

type usecase struct {
	hr repository.IHotpepperRepository
}

func NewUsecase(hr repository.IHotpepperRepository) IUsecase {
	return &usecase{hr}
}
