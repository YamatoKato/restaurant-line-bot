package igateway

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type IHotpepperGateway interface {
	GetRestaurantInfos(string) (*model.HotpepperResponse, error)
}
