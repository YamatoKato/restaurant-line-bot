package usecase

import "github.com/line/line-bot-sdk-go/v7/linebot"

func (u *usecase) SetAreaMenu() (*linebot.TemplateMessage, error) {
	bt := linebot.NewButtonsTemplate(
		"https://cdn.pixabay.com/photo/2017/08/17/07/47/travel-2650303_1280.jpg",
		"１/２　検索エリアを指定",
		"指定エリアまたは位置情報を送信し、その中心から検索します\n※エリアを指定する場合「---」を消さずに入力してください",
		linebot.NewPostbackAction("エリアを指定する", "area", "", "", "openKeyboard", "---\n東京 渋谷\n---"),
		linebot.NewURIAction("位置情報を送る", "https://line.me/R/nv/location/"),
	)
	tm := linebot.NewTemplateMessage("エリアを選択してください", bt)
	return tm, nil
}
