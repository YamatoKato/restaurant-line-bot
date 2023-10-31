package menudto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type SetConditionMenuInput struct {
	ReplyToken   string
	PostbackData model.PostbackData
}
