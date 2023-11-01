package model

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

// 次のアクションを識別する接頭辞
const (
	PBD_PREFIX_IDENTIFY_AREA      = "A"
	PBD_PREFIX_IDENTIFY_GENRE     = "G"
	PBD_PREFIX_IDENTIFY_CONDITION = "C"
	PBD_PREFIX_IDENTIFY_KEYWORD   = "K"
	PBD_PREFIX_IDENTIFY_BUDGET    = "B"
	PBD_PREFIX_IDENTIFY_SEARCH    = "S"
	PBD_PREFIX_IDENTIFY_CONFIRM   = "F"
)

// メッセージ
const (
	INTRO_MESSAGE                               = "お店を探す"
	PBA_DISPLAY_TEXT_GENRE_SET_CONFIRM_MENU     = "ジャンルを選択します\n\n以下からジャンルを選択してください"
	PBA_DISPLAY_TEXT_CONDITION_SET_CONFIRM_MENU = "条件を指定します\n\n以下から条件を選択してください"
	PBA_DISPLAY_TEXT_SET_AREA_MENU              = ""
	PBA_DISPLAY_TEXT_SET_CONFIRM_MENU           = ""
	PBA_DISPLAY_TEXT_SEARCH_SET_CONFIRM_MENU    = "検索します..."
	PBA_TEXT_SET_GENRE_MENU                     = ""
	PBA_DISPLAY_TEXT_SET_GENRE_MENU             = "「%s」\n\n※ジャンル選択を変更したい場合は、↑のジャンル一覧から再度選択してください"
	PBA_DISPLAY_TEXT_SET_CONDITION_MENU         = "「%s」を条件に追加しました\n\nさらに追加で条件を指定する場合は、↓から「%s」を選択してください\n\n※入力した条件を変更したい場合は、↑の条件一覧から再度選択してください"
	PBA_TEXT_SET_CONDITION_MENU                 = ""
	TM_CONTENT_SET_HELP_MENU                    = "ヘルプメニューです\n\nメニュー画面をタップして、お気に入りのエリアのお店を見つけてみましょう。簡単な対話形式でお店を検索できます。"
	TM_CONTENT_GET_RESTAURANT_INFOS_NOTHING_HIT = "該当するお店が見つかりませんでした。"
)

// サムネイル画像URL
const (
	BT_THUMBNAIL_SET_AREA_MENU    = "https://cdn.pixabay.com/photo/2017/08/17/07/47/travel-2650303_1280.jpg"
	BT_THUMBNAIL_SET_CONFIRM_MENU = ""
)

// アクションURI
const (
	UA_URI_SET_AREA_MENU = "https://line.me/R/nv/location/"
	GOOGLE_MAPS_URI      = "https://www.google.com/maps/search/?api=1&query=%s"
	TEL_URI              = "tel:%s"
)

// タイトル・ラベル
const (
	BT_TITLE_SET_AREA_MENU               = "検索エリアを指定"
	BT_TITLE_SET_CONFIRM_MENU            = ""
	PBA_LABEL_SET_AREA_MENU              = "エリアを指定する"
	TM_LABEL_SET_AREA_MENU               = "エリアを選択してください"
	TM_LABEL_SET_CONFIRM_MENU            = "検索条件を確定しますか"
	TM_LABEL_SET_GENRE_MENU              = "ジャンルを選択してください"
	TM_ALT_TEXT_SET_CONDITION_MENU       = "条件を選択してください"
	TM_ALT_TEXT_GET_RESTAURANT_INFOS     = "レストラン一覧"
	UA_LABEL_SET_AREA_MENU               = "位置情報を送る"
	PBA_LABEL_GENRE_SET_CONFIRM_MENU     = "ジャンルを選択する"
	PBA_LABEL_CONDITION_SET_CONFIRM_MENU = "追加で条件を指定する"
	PBA_LABEL_SEARCH_SET_CONFIRM_MENU    = "この条件で検索する"
	URI_LABEL_GET_RESTAURANT_INFOS_MORE  = "ホットペッパーで開く"
	URI_LABEL_GET_RESTAURANT_INFOS_MAP   = "マップを開く"
	URI_LABEL_GET_RESTAURANT_INFOS_TEL   = "電話する"
)

// カード内メッセージ等
const (
	BT_MESSAGE_SET_AREA_MENU  = "指定エリアまたは位置情報を送信し、その中心から検索します\n※エリアを指定する場合「---」を消さずに入力してください"
	PBA_TEXT_SET_AREA_MENU    = ""
	PBA_TEXT_SET_CONFIRM_MENU = ""

	PBA_INPUT_OPTION_SET_AREA_MENU      = "openKeyboard"
	PBA_FILL_IN_TEXT_SET_AREA_MENU      = "---\n東京 渋谷\n---"
	PBA_FILL_IN_TEXT_SET_CONFIRM_MENU   = ""
	PBA_FILL_IN_TEXT_SET_GENRE_MENU     = ""
	PBA_FILL_IN_TEXT_SET_CONDITION_MENU = ""

	CCM_TEXT_SET_GENRE_MENU       = "ジャンル「%s」\nエリア「%s」付近"
	CCM_TEXT_GET_RESTAURANT_INFOS = "住所「%s」\n■ %s\n■ %s"
)

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}

type PostbackData struct {
	AreaStr      string `json:"areaStr,omitempty"`
	Lat          string `json:"lat,omitempty"`
	Lng          string `json:"lng,omitempty"`
	GenreCode    string `json:"genreCode,omitempty"`
	Keyword      string `json:"keyword,omitempty"`
	Smoking      string `json:"smoking,omitempty"`
	Parking      string `json:"parking,omitempty"`
	PetFriendly  string `json:"pet,omitempty"`
	MidnightOpen string `json:"midnight,omitempty"`
	MidnightMeal string `json:"midnight_meal,omitempty"`
	PrivateRoom  string `json:"private_room,omitempty"`
	FreeFood     string `json:"free_food,omitempty"`
	FreeDrink    string `json:"free_drink,omitempty"`
	Budget       string `json:"budget,omitempty"`
	Terrace      string `json:"terrace,omitempty"`
}

type TemplateMessageData struct {
	AltText string
}

type TextMessageData struct {
	Content string
}

type ButtonsTemplateData struct {
	ThumbnailImageURL string
	Title             string
	Text              string
}

type PostbackActionData struct {
	Label       string
	Data        string
	Text        string
	DisplayText string
	InputOption linebot.InputOption
	FillInText  string
}

type URIActionData struct {
	Label string
	URI   string
}

type CarouselColumnData struct {
	ThumbnailImageURL string
	Title             string
	Text              string
}

func SetPostbackDataField(data PostbackData, fieldName string, value string) PostbackData {

	newData := data

	switch fieldName {
	case "areaStr":
		newData.AreaStr = value
	case "lat":
		newData.Lat = value
	case "lng":
		newData.Lng = value
	case "genreCode":
		newData.GenreCode = value
	case "keyword":
		newData.Keyword = value
	case "smoking":
		newData.Smoking = value
	case "parking":
		newData.Parking = value
	case "pet":
		newData.PetFriendly = value
	case "midnight":
		newData.MidnightOpen = value
	case "midnight_meal":
		newData.MidnightMeal = value
	case "private_room":
		newData.PrivateRoom = value
	case "free_food":
		newData.FreeFood = value
	case "free_drink":
		newData.FreeDrink = value
	case "budget":
		newData.Budget = value
	case "terrace":
		newData.Terrace = value
	default:
	}

	return newData
}

func GetPrefix(input string) string {
	if len(input) > 0 {
		switch input[0] {
		case 'A':
			return PBD_PREFIX_IDENTIFY_AREA
		case 'G':
			return PBD_PREFIX_IDENTIFY_GENRE
		case 'C':
			return PBD_PREFIX_IDENTIFY_CONDITION
		case 'K':
			return PBD_PREFIX_IDENTIFY_KEYWORD
		case 'B':
			return PBD_PREFIX_IDENTIFY_BUDGET
		case 'S':
			return PBD_PREFIX_IDENTIFY_SEARCH
		case 'F':
			return PBD_PREFIX_IDENTIFY_CONFIRM
		default:
			return ""
		}
	}
	return ""
}
