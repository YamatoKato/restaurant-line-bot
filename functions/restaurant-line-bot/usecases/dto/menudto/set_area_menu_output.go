package menudto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type SetAreaMenuOutput struct {
	ReplyToken          string
	ButtonsTemplateData model.ButtonsTemplateData
	PostbackActionData  model.PostbackActionData
	URIActionData       model.URIActionData
	TemplateMessageData model.TemplateMessageData
}
