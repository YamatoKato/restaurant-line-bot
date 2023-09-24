package main

import (
	"restaurant-line-bot/functions/restaurant-line-bot/infrastructure"

	"github.com/labstack/echo/v4"
)

func InitDI(e *echo.Echo) *infrastructure.Router {
	linebotController := controllers.NewLinebotController(searchInteractor)
	router := infrastructure.NewRouter(e, linebotController)
	return router
}
