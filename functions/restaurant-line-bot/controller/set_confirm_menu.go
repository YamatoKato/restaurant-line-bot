package controller

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func setConfirmMenu(c *Controller, e *linebot.Event, bot *linebot.Client, messageType string, postbackStr string, userMessage string) error {
	var data model.PostbackData

	if postbackStr == "" {
		areaWording := utils.GetAreaWording(userMessage)
		data = model.PostbackData{
			AreaStr: areaWording,
			Keyword: areaWording,
		}

		if userMessage == "" {
			msg := e.Message.(*linebot.LocationMessage)

			lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
			lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

			data = model.PostbackData{
				Lat:     lat,
				Lng:     lng,
				AreaStr: utils.GetAreaStrFromLocation(msg.Address),
			}
		}
	} else {
		if err := json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data); err != nil {
			fmt.Println(err, "controller@setConditionMenu_json.Unmarshal")
			return err
		}
	}

	res, err := c.u.SetConfirmMenu(&data, messageType)
	if err != nil {
		fmt.Println(err, "controller@setConfirmMenu")
		return err
	}

	if _, err := bot.ReplyMessage(e.ReplyToken, res).Do(); err != nil {
		fmt.Println(err, "controller@setConfirmMenu.ReplyMessage")
		return err
	}
	return nil

}
