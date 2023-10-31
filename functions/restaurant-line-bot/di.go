package main

import (
	"restaurant-line-bot/functions/restaurant-line-bot/infrastructure"
	"restaurant-line-bot/functions/restaurant-line-bot/interfaces/controllers"
	"restaurant-line-bot/functions/restaurant-line-bot/interfaces/gateway"
	"restaurant-line-bot/functions/restaurant-line-bot/interfaces/presenter"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/interactor"
)

func InitDI() *infrastructure.Router {
	hotpepperGateway := gateway.NewHotpepperGateway()
	linePresenter := presenter.NewLinePresenter()
	searchInteractor := interactor.NewSearchInteractor(hotpepperGateway, linePresenter)
	setMenuInteractor := interactor.NewSetMenuInteractor(linePresenter)
	linebotController := controllers.NewLinebotController(searchInteractor, setMenuInteractor)
	router := infrastructure.NewRouter(linebotController)
	return router
}
