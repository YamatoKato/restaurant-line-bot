package usecase

import (
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/searchdto"
)

type ISearchUsecase interface {
	GetRestaurantInfos(searchdto.Input) (searchdto.Output, error)
}
