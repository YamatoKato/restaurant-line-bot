package controller

import (
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func setGenreMenu(c *Controller, e *linebot.Event, bot *linebot.Client, userMessage string) error {
	areaWording := utils.GetAreaWording(userMessage)
	apiParams := &model.APIParams{
		Keyword: areaWording,
		AreaStr: areaWording,
	}

	if userMessage == "" {
		msg := e.Message.(*linebot.LocationMessage)

		lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
		lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

		apiParams = &model.APIParams{
			Lat:     lat,
			Lng:     lng,
			AreaStr: utils.GetAreaStrFromLocation(msg.Address),
		}
	}

	res, err := c.u.SetGenreMenu(apiParams)
	if err != nil {
		fmt.Println(err, "controller@setGenreMenu")
		return err
	}

	if _, err := bot.ReplyMessage(e.ReplyToken, res).Do(); err != nil {
		fmt.Println(err, "controller@HandleRequest_bot.ReplyMessage")
		return err
	}
	return nil
}
