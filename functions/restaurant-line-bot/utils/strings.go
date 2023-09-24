package utils

import (
	"encoding/json"
	"net/url"
	"os"
	"regexp"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"strings"
	"unicode/utf8"
)

// 正規表現で半角スペース、全角スペース、タブ、改行などの空白文字を取り除く
func RemoveSpaces(input string) string {
	re := regexp.MustCompile(`[\s　]+`)
	return re.ReplaceAllString(input, "")
}

// 指定文字数以上ある場合はそれ以降をカットし、「...」をつける
func CutString(text string, limit int) string {
	// 文字数を数える
	length := utf8.RuneCountInString(text)

	// 制限文字数以下ならそのまま返す
	if length <= limit {
		return text
	}

	// 制限文字数までカットして「...」を追加
	runes := []rune(text)
	return string(runes[:limit-3]) + "..."
}

// 正規表現を使用して「---」で囲まれたテキストを抽出
func GetAreaWording(input string) string {
	re := regexp.MustCompile("---\n(.*?)\n---")
	match := re.FindStringSubmatch(input)

	if len(match) >= 2 {
		area := strings.ReplaceAll(match[1], "　", " ")
		return area
	} else {
		return ""
	}
}

// 文字列に「---\n\n---」のテキストが含まれているかを判定
func ContainsHyphen(input string) bool {
	if strings.Contains(input, "---\n") && strings.Contains(input, "\n---") {
		return true
	} else {
		return false
	}
}

// APIのURLを構築する関数
func BuildAPIURL(apiParams model.PostbackData) string {
	// URLパラメーターを設定
	params := url.Values{}
	params.Add("format", "json")
	params.Add("key", os.Getenv("HOTPEPPER_API_KEY"))

	if apiParams.Lat != "" {
		params.Add("lat", apiParams.Lat)
	}

	if apiParams.Lng != "" {
		params.Add("lng", apiParams.Lng)
	}

	if apiParams.GenreCode != "" {
		params.Add("genre", apiParams.GenreCode)
	}

	if apiParams.Keyword != "" {
		params.Add("genre", apiParams.Keyword)
	}

	// URL文字列を組み立て
	apiURL := model.BASE_HOTPEPPER_API_URL + "?" + params.Encode()
	return apiURL
}

func CreatePostbackData(postbackData *model.PostbackData) (string, error) {
	jsonData, err := json.Marshal(postbackData)
	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func GetAreaStrFromLocation(input string) string {
	// 正規表現パターンを定義
	addressPattern := `日本、〒\d{3}-\d{4}\s(.*?)$`
	digitPattern := `[０-９]`

	// 正規表現をコンパイル
	addressRegex := regexp.MustCompile(addressPattern)
	digitRegex := regexp.MustCompile(digitPattern)

	// マッチした部分を抽出
	addressMatches := addressRegex.FindStringSubmatch(input)
	if len(addressMatches) >= 2 {
		// 最初に出てきた全角数字の位置を検出
		match := digitRegex.FindStringIndex(addressMatches[1])
		if match != nil {
			// 全角数字以降の部分を削除
			result := addressMatches[1][:match[0]]
			return result
		}
		return addressMatches[1] // 全角数字が見つからない場合、元の住所を返す
	}

	return ""
}

func CreateTextMessage(data model.PostbackData) string {
	baseStr := ""
	conditionStr := "\n\n追加した条件：\n"

	if data.AreaStr != "" {
		baseStr += "エリア：" + data.AreaStr + "\n"
	}
	if data.GenreCode != "" {
		baseStr += "ジャンル：" + model.SearchGenreNameByCode(data.GenreCode) + "\n"
	}
	if data.Keyword != "" {
		conditionStr += "キーワード：" + data.Keyword + "\n"
	}
	if data.Smoking != "" {
		conditionStr += "■" + " " + model.SMOKING_JP + "\n"
	}
	if data.Parking != "" {
		conditionStr += "■" + " " + model.PARKING_JP + "\n"
	}
	if data.PetFriendly != "" {
		conditionStr += "■" + " " + model.PET_FRIENDLY_JP + "\n"
	}
	if data.MidnightOpen != "" {
		conditionStr += "■" + " " + model.MIDNIGHT_OPEN_JP + "\n"
	}
	if data.MidnightMeal != "" {
		conditionStr += "■" + " " + model.MIDNIGHT_MEAL_JP + "\n"
	}
	if data.PrivateRoom != "" {
		conditionStr += "■" + " " + model.PRIVATE_ROOM_JP + "\n"
	}
	if data.Terrace != "" {
		conditionStr += "■" + " " + model.TERRACE_JP + "\n"
	}
	if data.FreeDrink != "" {
		conditionStr += "■" + " " + model.FREE_DRINK_JP + "\n"
	}
	if data.FreeFood != "" {
		conditionStr += "■" + " " + model.FREE_FOOD_JP + "\n"
	}

	return baseStr + conditionStr
}

func GetPrefix(input string) string {
	if len(input) > 0 {
		switch input[0] {
		case 'A':
			return model.PBD_PREFIX_IDENTIFY_AREA
		case 'G':
			return model.PBD_PREFIX_IDENTIFY_GENRE
		case 'C':
			return model.PBD_PREFIX_IDENTIFY_CONDITION
		case 'K':
			return model.PBD_PREFIX_IDENTIFY_KEYWORD
		case 'B':
			return model.PBD_PREFIX_IDENTIFY_BUDGET
		case 'S':
			return model.PBD_PREFIX_IDENTIFY_SEARCH
		case 'F':
			return model.PBD_PREFIX_IDENTIFY_CONFIRM
		default:
			return ""
		}
	}
	return ""
}

func RemoveFirstTwoCharacters(input string) string {
	// 文字列が2文字未満の場合は空文字列を返す
	if len(input) < 2 {
		return ""
	}

	// 先頭の2文字を削除して新しい文字列を作成する
	result := input[2:]

	return result
}
