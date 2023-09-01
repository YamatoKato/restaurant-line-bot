package model

import "github.com/line/line-bot-sdk-go/v7/linebot"

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}
