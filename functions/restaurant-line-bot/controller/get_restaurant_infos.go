package controller

import (
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetRestaurantInfos(c *Controller, e *linebot.Event, bot *linebot.Client) error {
	msg := e.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	area := model.Area{
		Latitude:  lat,
		Longitude: lng,
	}

	res, err := c.bu.GetRestaurantInfos(area)
	if err != nil {
		fmt.Println(err, "controller@getRestaurantInfos_bc.bu.GetRestaurantInfos")
		return err
	}

	if _, err := bot.ReplyMessage(e.ReplyToken, res).Do(); err != nil {
		fmt.Println(err, "controller@getRestaurantInfos_do")
		return err
	}

	return nil
}
