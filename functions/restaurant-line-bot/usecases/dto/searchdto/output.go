package searchdto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type Output struct {
	ReplyToken           string
	Response             *model.HotpepperResponse
	TemplateMessageData  *model.TemplateMessageData
	URIActionDataForMore *model.URIActionData
	URIActionDataForMap  *model.URIActionData
	TextMessageData      *model.TextMessageData
	// URIActionDataForTel  *model.URIActionData
}
