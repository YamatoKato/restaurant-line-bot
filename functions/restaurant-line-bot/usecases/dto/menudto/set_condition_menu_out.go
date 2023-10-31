package menudto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type SetConditionMenuOutput struct {
	ReplyToken          string
	Conditions          []model.Condition
	PostbackData        model.PostbackData
	TemplateMessageData model.TemplateMessageData
	PostbackActionData  model.PostbackActionData
}
