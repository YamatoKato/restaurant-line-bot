package usecase

import (
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/menudto"
)

type ISetMenuUsecase interface {
	SetHelpMenu(menudto.SetHelpMenuInput) (menudto.SetHelpMenuOutput, error)
	SetAreaMenu(menudto.SetAreaMenuInput) (menudto.SetAreaMenuOutput, error)
	SetConfirmMenu(menudto.SetConfirmMenuInput) (menudto.SetConfirmMenuOutput, error)
	SetGenreMenu(menudto.SetGenreMenuInput) (menudto.SetGenreMenuOutput, error)
	SetConditionMenu(menudto.SetConditionMenuInput) (menudto.SetConditionMenuOutput, error)
}
