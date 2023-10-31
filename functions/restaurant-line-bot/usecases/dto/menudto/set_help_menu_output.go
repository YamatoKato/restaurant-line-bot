package menudto

import "restaurant-line-bot/functions/restaurant-line-bot/domain/model"

type SetHelpMenuOutput struct {
	ReplyToken      string
	TextMessageData *model.TextMessageData
}
