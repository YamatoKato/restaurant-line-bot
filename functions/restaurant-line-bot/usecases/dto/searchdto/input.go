package searchdto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type Input struct {
	ReplyToken   string
	PostbackData model.PostbackData
}
