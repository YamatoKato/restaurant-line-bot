package controller

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func setGenreMenu(c *Controller, e *linebot.Event, bot *linebot.Client, postbackStr string) error {
	var data model.PostbackData
	if err := json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data); err != nil {
		fmt.Println(err, "controller@setGenreMenu_json.Unmarshal")
		return err
	}
	res, err := c.u.SetGenreMenu(&data)
	if err != nil {
		fmt.Println(err, "controller@setGenreMenu")
		return err
	}

	if _, err := bot.ReplyMessage(e.ReplyToken, res...).Do(); err != nil {
		fmt.Println(err, "controller@setGenreMenu.ReplyMessage")
		return err
	}
	return nil
}
