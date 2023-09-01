package controller

import (
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func GetRestaurantInfos(c *Controller, we *linebot.Event, bot *linebot.Client) error {
	msg := we.Message.(*linebot.LocationMessage)

	lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
	lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

	area := model.Area{
		Latitude:  lat,
		Longitude: lng,
	}

	ccs, err := c.bu.GetRestaurantInfos(area)
	if err != nil {
		fmt.Println(err, "controller@getRestaurantInfos_bc.bu.GetRestaurantInfos")
		return err
	}

	res := linebot.NewTemplateMessage(
		"レストラン一覧",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)
	if _, err := bot.ReplyMessage(we.ReplyToken, res).Do(); err != nil {
		fmt.Println(err, "controller@getRestaurantInfos_do")
		return err
	}

	return nil
}
