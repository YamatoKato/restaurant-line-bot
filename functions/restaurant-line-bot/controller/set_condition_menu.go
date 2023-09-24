package controller

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func setConditionMenu(c *Controller, e *linebot.Event, bot *linebot.Client, postbackStr string) error {

	var data model.PostbackData
	if err := json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data); err != nil {
		fmt.Println(err, "controller@setConditionMenu_json.Unmarshal")
		return err
	}

	res, err := c.u.SetConditionMenu(&data)
	if err != nil {
		fmt.Println(err, "controller@setConditionMenu")
		return err
	}

	if _, err := bot.ReplyMessage(e.ReplyToken, res...).Do(); err != nil {
		fmt.Println(err, "controller@setConditionMenu.ReplyMessage")
		return err
	}
	return nil
}
