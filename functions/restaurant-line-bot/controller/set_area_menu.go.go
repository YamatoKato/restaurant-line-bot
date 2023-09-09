package controller

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func setAreaMenu(c *Controller, e *linebot.Event, bot *linebot.Client) error {
	res, err := c.u.SetAreaMenu()
	if err != nil {
		fmt.Println(err, "controller@setAreaMenu")
		return err
	}

	if _, err := bot.ReplyMessage(e.ReplyToken, res).Do(); err != nil {
		fmt.Println(err, "controller@HandleRequest_bot.ReplyMessage")
		return err
	}
	return nil
}
