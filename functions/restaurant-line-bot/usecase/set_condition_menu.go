package usecase

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (u *usecase) SetConditionMenu(postbackData *model.PostbackData) ([]linebot.SendingMessage, error) {
	var firstCcs []*linebot.CarouselColumn
	var remainCcs []*linebot.CarouselColumn

	// var newPostbackData model.PostbackData

	conditions := []model.Condition{

		model.CreateCondition(model.PET_FRIENDLY_JP, model.PET_FRIENDLY_IMG_URL, model.PET_FRIENDLY_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.PET_FRIENDLY_PARAM_KEY, model.PET_FRIENDLY_PARAM_VALUE)),

		model.CreateCondition(model.TERRACE_JP, model.TERRACE_IMG_URL, model.TERRACE_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.TERRACE_PARAM_KEY, model.TERRACE_PARAM_VALUE)),

		model.CreateCondition(model.SMOKING_JP, model.SMOKING_IMG_URL, model.SMOKING_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.SMOKING_PARAM_KEY, model.SMOKING_PARAM_VALUE)),

		model.CreateCondition(model.PARKING_JP, model.PARKING_IMG_URL, model.PARKING_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.PARKING_PARAM_KEY, model.PARKING_PARAM_VALUE)),

		model.CreateCondition(model.MIDNIGHT_OPEN_JP, model.MIDNIGHT_IMG_URL, model.MIDNIGHT_OPEN_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.MIDNIGHT_OPEN_PARAM_KEY, model.MIDNIGHT_OPEN_PARAM_VALUE)),

		model.CreateCondition(model.MIDNIGHT_MEAL_JP, model.MIDNIGHT_IMG_URL, model.MIDNIGHT_MEAL_DESC, model.CreateConditionOption(model.ADD_CONDITION_WORDING, model.MIDNIGHT_MEAL_PARAM_KEY, model.MIDNIGHT_MEAL_PARAM_VALUE)),

		model.CreateCondition(model.FREE_DRINK_JP, model.FREE_DRINK_IMG_URL, model.FREE_DRINK_DESC, model.CreateConditionOption(model.FREE_DRINK_JP, model.FREE_DRINK_PARAM_KEY, model.FREE_DRINK_PARAM_VALUE)),

		model.CreateCondition(model.FREE_FOOD_JP, model.FREE_FOOD_IMG_URL, model.FREE_DESC, model.CreateConditionOption(model.FREE_FOOD_JP, model.FREE_FOOD_PARAM_KEY, model.FREE_FOOD_PARAM_VALUE)),

		model.CreateCondition(model.PRIVATE_ROOM_JP, model.PRIVATE_ROOM_IMG_URL, model.PRIVATE_ROOM_DESC, model.CreateConditionOption(model.PRIVATE_ROOM_JP, model.PRIVATE_ROOM_PARAM_KEY, model.PRIVATE_ROOM_PARAM_VALUE)),

		model.CreateCondition(model.NON_CONDITION_WORDING, model.NON_CONDITION_IMG_URL, model.NON_CONDITION_DESC, model.CreateConditionOption(model.NON_CONDITION_WORDING, "", "")),

		model.CreateCondition(model.BUDGET_JP, model.BUDGET_IMG_URL, model.BUDGET_DESC, model.CreateConditionOption(model.ADD_CONDITION_BUDGET_WORDING, model.PBD_PREFIX_IDENTIFY_BUDGET, model.BUDGET_PARAM_VALUE)),

		model.CreateCondition(model.INPUT_JP, model.INPUT_IMG_URL, model.INPUT_DESC, model.CreateConditionOption(model.ADD_CONDITION_INPUT_WORDING, model.PBD_PREFIX_IDENTIFY_KEYWORD, model.INPUT_PARAM_VALUE)),
	}

	// 最小の10個までを表示
	for _, condition := range conditions[:10] {
		actions := []linebot.TemplateAction{}
		for _, option := range condition.ConditionOptions {
			newPostbackData := model.SetPostbackDataField(*postbackData, option.ParamKey, option.ParamValue)
			jsonData, err := json.Marshal(newPostbackData)
			if err != nil {
				fmt.Println(err, "usecase.set_condition_menu.json.Marshal(postbackData)")
				return nil, err
			}
			// ConditionOptionごとにアクションを作成
			action := linebot.NewPostbackAction(option.Name, model.PBD_PREFIX_IDENTIFY_CONFIRM+" "+string(jsonData), "", fmt.Sprintf("「%s」を条件に追加しました\n\nさらに追加で条件を指定する場合は、↓から「%s」を選択してください\n\n※入力した条件を変更したい場合は、↑の条件一覧から再度選択してください", condition.Name, model.BUTTON_MESSAGE_CONFIRM_ADD_CONDITION), "", "")
			actions = append(actions, action)
		}

		cc := linebot.NewCarouselColumn(
			condition.ImgURL,
			condition.Name,
			condition.Desc,
			actions...,
		).WithImageOptions("#FFFFFF")
		firstCcs = append(firstCcs, cc)
	}

	// 11個目以降を表示
	for _, condition := range conditions[10:] {
		actions := []linebot.TemplateAction{}
		for _, option := range condition.ConditionOptions {
			// ConditionOptionごとにアクションを作成
			action := linebot.NewPostbackAction(option.Name, option.ParamKey, "", option.Name, "", "")
			actions = append(actions, action)
		}

		cc := linebot.NewCarouselColumn(
			condition.ImgURL,
			condition.Name,
			condition.Desc,
			actions...,
		).WithImageOptions("#FFFFFF")
		remainCcs = append(remainCcs, cc)
	}

	var sms []linebot.SendingMessage

	firstTm := linebot.NewTemplateMessage("条件を選択してください", linebot.NewCarouselTemplate(firstCcs...).WithImageOptions("rectangle", "cover"))

	remainTm := linebot.NewTemplateMessage("条件を選択してください", linebot.NewCarouselTemplate(remainCcs...).WithImageOptions("rectangle", "cover"))

	sms = append(sms, firstTm, remainTm)

	return sms, nil
}
