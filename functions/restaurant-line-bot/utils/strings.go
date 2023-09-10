package utils

import (
	"fmt"
	"net/url"
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

// 指定文字数以上ある場合はそれ以降をカット
func CutString(input string, num int64) string {
	if 60 < utf8.RuneCountInString(input) {
		input = string([]rune(input)[:num])
	}
	return input
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
func BuildAPIURL(apiParams *model.APIParams) string {
	// URLパラメーターを設定
	params := url.Values{}
	params.Add("format", "json")
	params.Add("key", apiParams.APIKey)

	if apiParams.Lat != "" {
		params.Add("lat", apiParams.Lat)
	}

	if apiParams.Lng != "" {
		params.Add("lng", apiParams.Lng)
	}

	if apiParams.Genre != "" {
		params.Add("genre", apiParams.Genre)
	}

	if apiParams.Keyword != "" {
		params.Add("genre", apiParams.Keyword)
	}

	// URL文字列を組み立て
	apiURL := model.BASE_HOTPEPPER_API_URL + "?" + params.Encode()
	return apiURL
}

func CreatePostBackData(apiParams *model.APIParams, genreCode string) string {
	postBackData := ""

	if apiParams.Keyword != "" {
		postBackData = fmt.Sprintf("area=%s&lat=nil&lng=nil&genreCode=%s", apiParams.Keyword, genreCode)
	} else {
		postBackData = fmt.Sprintf("area=nil&lat=%s&lng=%s&genreCode=%s", apiParams.Lat, apiParams.Lng, genreCode)
	}

	return postBackData
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

func SearchGenreNameByCode(genreCode string) string {
	switch genreCode {
	case model.GENRE_JAPANESE_CODE:
		return model.GENRE_JAPANESE_JP
	case model.GENRE_WESTERN_CODE:
		return model.GENRE_WESTERN_JP
	case model.GENRE_CHINESE_CODE:
		return model.GENRE_CHINESE_JP
	case model.GENRE_ITALIAN_FRENCH_CODE:
		return model.GENRE_ITALIAN_FRENCH_JP
	case model.GENRE_KOREAN_CODE:
		return model.GENRE_KOREAN_JP
	case model.GENRE_RAMEN_CODE:
		return model.GENRE_RAMEN_JP
	case model.GENRE_YAKINIKU_OFFAL_CODE:
		return model.GENRE_YAKINIKU_OFFAL_JP
	case model.GENRE_CAFE_SWEETS_CODE:
		return model.GENRE_CAFE_SWEETS_JP
	case model.GENRE_IZAKAYA_CODE:
		return model.GENRE_IZAKAYA_JP
	case model.GENRE_ROULETTE_CODE:
		return model.GENRE_ROULETTE_JP
	// 他のジャンルに対する定義を追加
	default:
		return ""
	}
}

func SearchGenreImgUrlByCode(genreCode string) string {
	switch genreCode {
	case model.GENRE_JAPANESE_CODE:
		return model.GENRE_JAPANESE_IMG_URL
	case model.GENRE_WESTERN_CODE:
		return model.GENRE_WESTERN_IMG_URL
	case model.GENRE_CHINESE_CODE:
		return model.GENRE_CHINESE_IMG_URL
	case model.GENRE_ITALIAN_FRENCH_CODE:
		return model.GENRE_ITALIAN_FRENCH_IMG_URL
	case model.GENRE_KOREAN_CODE:
		return model.GENRE_KOREAN_IMG_URL
	case model.GENRE_RAMEN_CODE:
		return model.GENRE_RAMEN_IMG_URL
	case model.GENRE_YAKINIKU_OFFAL_CODE:
		return model.GENRE_YAKINIKU_OFFAL_IMG_URL
	case model.GENRE_CAFE_SWEETS_CODE:
		return model.GENRE_CAFE_SWEETS_IMG_URL
	case model.GENRE_IZAKAYA_CODE:
		return model.GENRE_IZAKAYA_IMG_URL
	// 他のジャンルに対する定義を追加
	case model.GENRE_ROULETTE_CODE:
		return model.GENRE_ROULETTE_IMG_URL
	default:
		return ""
	}
}
