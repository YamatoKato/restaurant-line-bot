package menudto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type SetConfirmMenuInput struct {
	ReplyToken   string
	AreaMessage  string
	PostbackData model.PostbackData
	MessageType  string
}
