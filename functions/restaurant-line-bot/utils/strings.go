package utils

import (
	"encoding/json"
	"regexp"
	"restaurant-line-bot/functions/restaurant-line-bot/domain/model"
	"strings"
	"unicode/utf8"
)

// 正規表現で半角スペース、全角スペース、タブ、改行などの空白文字を取り除く
func RemoveSpaces(input string) string {
	re := regexp.MustCompile(`[\s　]+`)
	return re.ReplaceAllString(input, "")
}

// 指定文字数以上ある場合はそれ以降をカットし、「...」をつける
func CutString(input string, length int) string {
	if utf8.RuneCountInString(input) > length {
		return string([]rune(input)[:length]) + "..."
	} else {
		return input
	}
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

func RemoveFirstTwoCharacters(input string) string {
	// 文字列が2文字未満の場合は空文字列を返す
	if len(input) < 2 {
		return ""
	}

	// 先頭の2文字を削除して新しい文字列を作成する
	result := input[2:]

	return result
}
