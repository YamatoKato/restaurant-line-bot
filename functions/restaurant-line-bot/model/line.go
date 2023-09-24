package model

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

const (
	INTRO_MESSAGE                        = "お店を探す"
	PBD_PREFIX_IDENTIFY_AREA             = "A"
	PBD_PREFIX_IDENTIFY_GENRE            = "G"
	PBD_PREFIX_IDENTIFY_CONDITION        = "C"
	PBD_PREFIX_IDENTIFY_KEYWORD          = "K"
	PBD_PREFIX_IDENTIFY_BUDGET           = "B"
	PBD_PREFIX_IDENTIFY_SEARCH           = "S"
	PBD_PREFIX_IDENTIFY_CONFIRM          = "F"
	BUTTON_MESSAGE_CONFIRM_ADD_CONDITION = "追加で条件を指定する"
	BUTTON_MESSAGE_CONFIRM_SELECT_GENRE  = "ジャンルを選択する"
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

func SetPostbackDataField(data PostbackData, fieldName string, value string) PostbackData {

	newData := data

	switch fieldName {
	case "areaStr":
		fmt.Println("areaStr:", value)
		newData.AreaStr = value
	case "lat":
		fmt.Println("lat:", value)
		newData.Lat = value
	case "lng":
		fmt.Println("lng:", value)
		newData.Lng = value
	case "genreCode":
		fmt.Println("genreCode:", value)
		newData.GenreCode = value
	case "keyword":
		fmt.Println("keyword:", value)
		newData.Keyword = value
	case "smoking":
		fmt.Println("smoking:", value)
		newData.Smoking = value
	case "parking":
		fmt.Println("parking:", value)
		newData.Parking = value
	case "pet":
		fmt.Println("pet:", value)
		newData.PetFriendly = value
	case "midnight":
		fmt.Println("midnight:", value)
		newData.MidnightOpen = value
	case "midnight_meal":
		fmt.Println("midnight_meal:", value)
		newData.MidnightMeal = value
	case "private_room":
		fmt.Println("private_room:", value)
		newData.PrivateRoom = value
	case "free_food":
		fmt.Println("free_food:", value)
		newData.FreeFood = value
	case "free_drink":
		fmt.Println("free_drink:", value)
		newData.FreeDrink = value
	case "budget":
		fmt.Println("budget:", value)
		newData.Budget = value
	case "terrace":
		fmt.Println("terrace:", value)
		newData.Terrace = value
	default:
		fmt.Println("無効なフィールド名:", fieldName)
	}

	return newData
}

func GetTypeByPBDPrefix(prefix string) string {
	switch prefix {
	case PBD_PREFIX_IDENTIFY_AREA:
		return "area"
	case PBD_PREFIX_IDENTIFY_GENRE:
		return "genre"
	case PBD_PREFIX_IDENTIFY_CONDITION:
		return "condition"
	case PBD_PREFIX_IDENTIFY_KEYWORD:
		return "keyword"
	case PBD_PREFIX_IDENTIFY_BUDGET:
		return "budget"
	case PBD_PREFIX_IDENTIFY_SEARCH:
		return "search"
	case PBD_PREFIX_IDENTIFY_CONFIRM:
		return "confirm"
	default:
		return ""
	}
}
