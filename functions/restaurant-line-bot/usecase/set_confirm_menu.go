package usecase

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (u *usecase) SetConfirmMenu(postbackData *model.PostbackData, messageType string) (*linebot.TemplateMessage, error) {
	label := ""
	prefix := ""
	displayText := ""
	if messageType == model.PBD_PREFIX_IDENTIFY_GENRE {
		label = model.BUTTON_MESSAGE_CONFIRM_SELECT_GENRE
		displayText = "ジャンルを選択します\n\n以下からジャンルを選択してください"
		prefix = model.PBD_PREFIX_IDENTIFY_GENRE
	} else {
		label = model.BUTTON_MESSAGE_CONFIRM_ADD_CONDITION
		displayText = "条件を指定します\n\n以下から条件を選択してください"
		prefix = model.PBD_PREFIX_IDENTIFY_CONDITION
	}

	jsonData, err := json.Marshal(postbackData)
	if err != nil {
		fmt.Println(err, "usecase.set_confirm_menu.json.Marshal(postbackData)")
		return nil, err
	}
	bt := linebot.NewButtonsTemplate(
		"",
		"",
		utils.CreateTextMessage(*postbackData),
		linebot.NewPostbackAction(label, prefix+" "+string(jsonData), "", displayText, "", ""),
		linebot.NewPostbackAction("この条件で検索する", model.PBD_PREFIX_IDENTIFY_SEARCH+" "+string(jsonData), "", "検索します...", "", ""),
	)

	tm := linebot.NewTemplateMessage("検索条件を確定しますか？", bt)

	return tm, nil
}
