package utils

import (
	"regexp"
	"unicode/utf8"
)

func RemoveSpaces(input string) string {
	// 正規表現で半角スペース、全角スペース、タブ、改行などの空白文字を検索
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
