package model

import "github.com/line/line-bot-sdk-go/v7/linebot"

const (
	INTRO_MESSAGE = "お店を探す"
)

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}
