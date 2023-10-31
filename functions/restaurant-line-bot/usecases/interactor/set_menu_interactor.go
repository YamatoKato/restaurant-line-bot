package interactor

import (
	"restaurant-line-bot/functions/restaurant-line-bot/domain/model"

	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/menudto"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/ipresenter"

	"github.com/sirupsen/logrus"
)

// SetMenuInteractor メニューインタラクタ
type SetMenuInteractor struct {
	linePresenter ipresenter.ILinePresenter
}

// NewSetMenuInteractor コンストラクタ
func NewSetMenuInteractor(
	linePresenter ipresenter.ILinePresenter) *SetMenuInteractor {
	return &SetMenuInteractor{
		linePresenter: linePresenter,
	}
}

func (interactor *SetMenuInteractor) SetHelpMenu(in menudto.SetHelpMenuInput) (menudto.SetHelpMenuOutput, error) {
	out := menudto.SetHelpMenuOutput{
		ReplyToken: in.ReplyToken,
		TextMessageData: &model.TextMessageData{
			Content: model.TM_CONTENT_SET_HELP_MENU,
		},
	}

	if err := interactor.linePresenter.SetHelpMenuReplyMessage(out); err != nil {
		logrus.Errorf("Error linePresenter.SetHelpMenuReplyMessage(out): %v", err)
		return out, err
	}
	return out, nil
}

// SetAreaMenu エリアメニュー設定
func (interactor *SetMenuInteractor) SetAreaMenu(in menudto.SetAreaMenuInput) (menudto.SetAreaMenuOutput, error) {
	postbackData := "area_menu"

	out := menudto.SetAreaMenuOutput{
		ReplyToken: in.ReplyToken,
		ButtonsTemplateData: model.ButtonsTemplateData{
			ThumbnailImageURL: model.BT_THUMBNAIL_SET_AREA_MENU,
			Title:             model.BT_TITLE_SET_AREA_MENU,
			Text:              model.BT_MESSAGE_SET_AREA_MENU,
		},
		PostbackActionData: model.PostbackActionData{
			Label:       model.PBA_LABEL_SET_AREA_MENU,
			Data:        postbackData,
			Text:        model.PBA_TEXT_SET_AREA_MENU,
			DisplayText: model.PBA_DISPLAY_TEXT_SET_AREA_MENU,
			InputOption: model.PBA_INPUT_OPTION_SET_AREA_MENU,
			FillInText:  model.PBA_FILL_IN_TEXT_SET_AREA_MENU,
		},
		URIActionData: model.URIActionData{
			Label: model.UA_LABEL_SET_AREA_MENU,
			URI:   model.UA_URI_SET_AREA_MENU,
		},
		TemplateMessageData: model.TemplateMessageData{
			AltText: model.TM_LABEL_SET_AREA_MENU,
		},
	}
	if err := interactor.linePresenter.SetAreaMenuReplyMessage(out); err != nil {
		return out, err
	}
	return out, nil
}

// SetConfirmMenu 確認メニュー設定
func (interactor *SetMenuInteractor) SetConfirmMenu(in menudto.SetConfirmMenuInput) (menudto.SetConfirmMenuOutput, error) {
	label := ""
	prefix := ""
	displayText := ""

	if in.MessageType == model.PBD_PREFIX_IDENTIFY_GENRE {
		// ジャンルメニューに遷移
		label = model.PBA_LABEL_GENRE_SET_CONFIRM_MENU
		displayText = model.PBA_DISPLAY_TEXT_GENRE_SET_CONFIRM_MENU
		prefix = model.PBD_PREFIX_IDENTIFY_GENRE
	} else {
		// 条件メニューに遷移
		label = model.PBA_LABEL_CONDITION_SET_CONFIRM_MENU
		displayText = model.PBA_DISPLAY_TEXT_CONDITION_SET_CONFIRM_MENU
		prefix = model.PBD_PREFIX_IDENTIFY_CONDITION
	}

	out := menudto.SetConfirmMenuOutput{
		ReplyToken: in.ReplyToken,
		ButtonsTemplateData: model.ButtonsTemplateData{
			ThumbnailImageURL: model.BT_THUMBNAIL_SET_CONFIRM_MENU,
			Title:             model.BT_TITLE_SET_CONFIRM_MENU,
			Text:              model.CreateTextMessage(in.PostbackData),
		},
		PostbackActionData: model.PostbackActionData{
			Label:       label,
			Data:        prefix + " ",
			Text:        model.PBA_TEXT_SET_CONFIRM_MENU,
			DisplayText: displayText,
			FillInText:  model.PBA_FILL_IN_TEXT_SET_CONFIRM_MENU,
		},
		PostbackActionDataForSearch: model.PostbackActionData{
			Label:       model.PBA_LABEL_SEARCH_SET_CONFIRM_MENU,
			Data:        model.PBD_PREFIX_IDENTIFY_SEARCH + " ",
			Text:        model.PBA_TEXT_SET_CONFIRM_MENU,
			DisplayText: model.PBA_DISPLAY_TEXT_SEARCH_SET_CONFIRM_MENU,
			FillInText:  model.PBA_FILL_IN_TEXT_SET_CONFIRM_MENU,
		},
		TemplateMessageData: model.TemplateMessageData{
			AltText: model.TM_LABEL_SET_CONFIRM_MENU,
		},
		PostbackData: in.PostbackData,
	}

	if err := interactor.linePresenter.SetConfirmMenuReplyMessage(out); err != nil {
		return out, err
	}

	return out, nil

}

func (interactor *SetMenuInteractor) SetGenreMenu(in menudto.SetGenreMenuInput) (menudto.SetGenreMenuOutput, error) {
	genres := []model.Genre{
		model.CreateGenre(model.GENRE_JAPANESE_CODE, model.GENRE_JAPANESE_JP, model.GENRE_JAPANESE_IMG_URL),
		model.CreateGenre(model.GENRE_WESTERN_CODE, model.GENRE_WESTERN_JP, model.GENRE_WESTERN_IMG_URL),
		model.CreateGenre(model.GENRE_CHINESE_CODE, model.GENRE_CHINESE_JP, model.GENRE_CHINESE_IMG_URL),
		model.CreateGenre(model.GENRE_ITALIAN_FRENCH_CODE, model.GENRE_ITALIAN_FRENCH_JP, model.GENRE_ITALIAN_FRENCH_IMG_URL),
		model.CreateGenre(model.GENRE_KOREAN_CODE, model.GENRE_KOREAN_JP, model.GENRE_KOREAN_IMG_URL),
		model.CreateGenre(model.GENRE_RAMEN_CODE, model.GENRE_RAMEN_JP, model.GENRE_RAMEN_IMG_URL),
		model.CreateGenre(model.GENRE_YAKINIKU_OFFAL_CODE, model.GENRE_YAKINIKU_OFFAL_JP, model.GENRE_YAKINIKU_OFFAL_IMG_URL),
		model.CreateGenre(model.GENRE_CAFE_SWEETS_CODE, model.GENRE_CAFE_SWEETS_JP, model.GENRE_CAFE_SWEETS_IMG_URL),
		model.CreateGenre(model.GENRE_IZAKAYA_CODE, model.GENRE_IZAKAYA_JP, model.GENRE_IZAKAYA_IMG_URL),
		model.CreateGenre(model.GENRE_LOOK_MORE_CODE, model.GENRE_LOOK_MORE_JP, model.GENRE_LOOK_MORE_IMG_URL),
	}

	output := menudto.SetGenreMenuOutput{
		ReplyToken: in.ReplyToken,
		TemplateMessageData: model.TemplateMessageData{
			AltText: model.TM_LABEL_SET_GENRE_MENU,
		},
		AreaStr:      in.PostbackData.AreaStr,
		Genres:       genres,
		PostbackData: in.PostbackData,
		PostbackActionData: model.PostbackActionData{
			Text:        model.PBA_TEXT_SET_GENRE_MENU,
			Data:        model.PBD_PREFIX_IDENTIFY_CONFIRM + " ",
			DisplayText: model.PBA_DISPLAY_TEXT_SET_GENRE_MENU,
			FillInText:  model.PBA_FILL_IN_TEXT_SET_GENRE_MENU,
		},
		CarouselColumnData: model.CarouselColumnData{
			Text: model.CCM_TEXT_SET_GENRE_MENU,
		},
	}

	if err := interactor.linePresenter.SetGenreMenuReplyMessage(output); err != nil {
		return output, err
	}

	return output, nil

}

func (interactor *SetMenuInteractor) SetConditionMenu(in menudto.SetConditionMenuInput) (menudto.SetConditionMenuOutput, error) {
	conditions := []model.Condition{

		model.CreateCondition(model.PET_FRIENDLY_JP, model.PET_FRIENDLY_IMG_URL, model.PET_FRIENDLY_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.PET_FRIENDLY_PARAM_KEY, model.PET_FRIENDLY_PARAM_VALUE)),

		model.CreateCondition(model.TERRACE_JP, model.TERRACE_IMG_URL, model.TERRACE_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.TERRACE_PARAM_KEY, model.TERRACE_PARAM_VALUE)),

		model.CreateCondition(model.SMOKING_JP, model.SMOKING_IMG_URL, model.SMOKING_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.SMOKING_PARAM_KEY, model.SMOKING_PARAM_VALUE)),

		model.CreateCondition(model.PARKING_JP, model.PARKING_IMG_URL, model.PARKING_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.PARKING_PARAM_KEY, model.PARKING_PARAM_VALUE)),

		model.CreateCondition(model.MIDNIGHT_OPEN_JP, model.MIDNIGHT_IMG_URL, model.MIDNIGHT_OPEN_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.MIDNIGHT_OPEN_PARAM_KEY, model.MIDNIGHT_OPEN_PARAM_VALUE)),

		model.CreateCondition(model.MIDNIGHT_MEAL_JP, model.MIDNIGHT_IMG_URL, model.MIDNIGHT_MEAL_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.MIDNIGHT_MEAL_PARAM_KEY, model.MIDNIGHT_MEAL_PARAM_VALUE)),

		model.CreateCondition(model.FREE_DRINK_JP, model.FREE_DRINK_IMG_URL, model.FREE_DRINK_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.FREE_DRINK_PARAM_KEY, model.FREE_DRINK_PARAM_VALUE)),

		model.CreateCondition(model.FREE_FOOD_JP, model.FREE_FOOD_IMG_URL, model.FREE_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.FREE_FOOD_PARAM_KEY, model.FREE_FOOD_PARAM_VALUE)),

		model.CreateCondition(model.PRIVATE_ROOM_JP, model.PRIVATE_ROOM_IMG_URL, model.PRIVATE_ROOM_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.PRIVATE_ROOM_PARAM_KEY, model.PRIVATE_ROOM_PARAM_VALUE)),

		model.CreateCondition(model.NON_CONDITION_WORDING, model.NON_CONDITION_IMG_URL, model.NON_CONDITION_DESC, model.CreateConditionOption(model.NON_CONDITION_WORDING, "", "")),

		// model.CreateCondition(model.BUDGET_JP, model.BUDGET_IMG_URL, model.BUDGET_DESC, model.CreateConditionOption(model.ADD_CONDITION_BUDGET_WORDING, model.PBD_PREFIX_IDENTIFY_BUDGET, model.BUDGET_PARAM_VALUE)),

		// model.CreateCondition(model.INPUT_JP, model.INPUT_IMG_URL, model.INPUT_DESC, model.CreateConditionOption(model.ADD_CONDITION_INPUT_WORDING, model.PBD_PREFIX_IDENTIFY_KEYWORD, model.INPUT_PARAM_VALUE)),
	}

	output := menudto.SetConditionMenuOutput{
		ReplyToken: in.ReplyToken,
		TemplateMessageData: model.TemplateMessageData{
			AltText: model.TM_ALT_TEXT_SET_CONDITION_MENU,
		},
		PostbackActionData: model.PostbackActionData{
			Data:        model.PBD_PREFIX_IDENTIFY_CONFIRM + " ",
			Text:        model.PBA_TEXT_SET_CONDITION_MENU,
			DisplayText: model.PBA_DISPLAY_TEXT_SET_CONDITION_MENU,
			FillInText:  model.PBA_FILL_IN_TEXT_SET_CONDITION_MENU,
		},
		Conditions:   conditions,
		PostbackData: in.PostbackData,
	}

	if err := interactor.linePresenter.SetConditionMenuReplyMessage(output); err != nil {
		return output, err
	}

	return output, nil
}
