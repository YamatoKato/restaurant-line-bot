package ipresenter

import (
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/menudto"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/searchdto"
)

type ILinePresenter interface {
	SetHelpMenuReplyMessage(menudto.SetHelpMenuOutput) error
	SetAreaMenuReplyMessage(menudto.SetAreaMenuOutput) error
	SetConfirmMenuReplyMessage(menudto.SetConfirmMenuOutput) error
	SetGenreMenuReplyMessage(menudto.SetGenreMenuOutput) error
	SetConditionMenuReplyMessage(menudto.SetConditionMenuOutput) error
	SetRestaurantListReplyMessage(searchdto.Output) error
	SetNothingHitReplyMessage(searchdto.Output) error
}
