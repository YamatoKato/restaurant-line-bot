package menudto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type SetGenreMenuOutput struct {
	ReplyToken          string
	TemplateMessageData model.TemplateMessageData
	AreaStr             string
	Genres              []model.Genre
	PostbackData        model.PostbackData
	PostbackActionData  model.PostbackActionData
	CarouselColumnData  model.CarouselColumnData
}
