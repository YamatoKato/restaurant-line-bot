package menudto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type SetConfirmMenuOutput struct {
	ReplyToken                  string
	ButtonsTemplateData         model.ButtonsTemplateData
	PostbackActionData          model.PostbackActionData
	PostbackActionDataForSearch model.PostbackActionData
	TemplateMessageData         model.TemplateMessageData
	PostbackData                model.PostbackData
}
