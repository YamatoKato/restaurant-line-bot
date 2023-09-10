package model

const BASE_HOTPEPPER_API_URL = "https://webservice.recruit.co.jp/hotpepper/gourmet/v1/"

// 各ジャンルに対する定数を定義
const (
	GENRE_IZAKAYA_CODE           = "G001" // 居酒屋
	GENRE_IZAKAYA_JP             = "居酒屋"
	GENRE_IZAKAYA_IMG_URL        = "https://cdn.pixabay.com/photo/2017/08/13/08/19/cheers-2636510_1280.jpg"
	GENRE_DINING_BAR_BAR         = "G002" // ダイニングバー・バル
	GENRE_CREATIVE_CUISINE       = "G003" // 創作料理
	GENRE_JAPANESE_CODE          = "G004" // 和食
	GENRE_JAPANESE_JP            = "和食"
	GENRE_JAPANESE_IMG_URL       = "https://cdn.pixabay.com/photo/2016/08/19/09/30/japan-1604865_1280.jpg"
	GENRE_WESTERN_CODE           = "G005" // 洋食
	GENRE_WESTERN_JP             = "洋食"
	GENRE_WESTERN_IMG_URL        = "https://cdn.pixabay.com/photo/2018/09/11/16/16/food-3669920_1280.jpg"
	GENRE_ITALIAN_FRENCH_CODE    = "G006" // イタリアン・フレンチ
	GENRE_ITALIAN_FRENCH_JP      = "イタリアン・フレンチ"
	GENRE_ITALIAN_FRENCH_IMG_URL = "https://cdn.pixabay.com/photo/2023/01/17/07/59/mossel-dish-7724002_1280.jpg"
	GENRE_CHINESE_CODE           = "G007" // 中華
	GENRE_CHINESE_JP             = "中華"
	GENRE_CHINESE_IMG_URL        = "https://cdn.pixabay.com/photo/2017/05/26/13/59/dim-sum-2346105_1280.jpg"
	GENRE_YAKINIKU_OFFAL_CODE    = "G008" // 焼肉・ホルモン
	GENRE_YAKINIKU_OFFAL_JP      = "焼肉・ホルモン"
	GENRE_YAKINIKU_OFFAL_IMG_URL = "https://cdn.pixabay.com/photo/2015/02/13/11/22/hokkaido-635019_1280.jpg"
	GENRE_KOREAN_CODE            = "G017" // 韓国料理
	GENRE_KOREAN_JP              = "韓国料理"
	GENRE_KOREAN_IMG_URL         = "https://cdn.pixabay.com/photo/2016/07/22/05/07/delicious-1534207_1280.jpg"
	GENRE_ASIA_ETHNIC_CUISINE    = "G009" // アジア・エスニック料理
	GENRE_INTERNATIONAL_CUISINE  = "G010" // 各国料理
	GENRE_KARAOKE_PARTY          = "G011" // カラオケ・パーティ
	GENRE_BAR_COCKTAIL           = "G012" // バー・カクテル
	GENRE_RAMEN_CODE             = "G013" // ラーメン
	GENRE_RAMEN_JP               = "ラーメン"
	GENRE_RAMEN_IMG_URL          = "https://cdn.pixabay.com/photo/2019/11/23/15/26/ramen-4647408_1280.jpg"
	GENRE_OKONOMIYAKI_MONJA      = "G016" // お好み焼き・もんじゃ
	GENRE_CAFE_SWEETS_CODE       = "G014" // カフェ・スイーツ
	GENRE_CAFE_SWEETS_JP         = "カフェ・スイーツ"
	GENRE_CAFE_SWEETS_IMG_URL    = "https://cdn.pixabay.com/photo/2016/11/29/12/54/cafe-1869656_1280.jpg"
	GENRE_OTHER_GOURMET          = "G015" // その他グルメ
	GENRE_ROULETTE_CODE          = "G999" // ルーレット
	GENRE_ROULETTE_JP            = "おまかせ1件"
	GENRE_ROULETTE_IMG_URL       = "https://cdn.pixabay.com/photo/2021/07/20/08/51/roulette-6480112_1280.jpg"
)

type Area struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

type APIParams struct {
	APIKey  string
	Lat     string
	Lng     string
	Genre   string
	Keyword string
	AreaStr string
}

type Genre struct {
	Code   string
	Name   string
	ImgURL string
}

// response APIレスポンス
type HotpepperResponse struct {
	Results results `json:"results"`
}

// results APIレスポンスの内容
type results struct {
	Shop []shop `json:"shop"`
}

// shop レストラン一覧
type shop struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Photo   photo  `json:"photo"`
	URLS    urls   `json:"urls"`
}

// photo 写真URL一覧
type photo struct {
	Mobile mobile `json:"mobile"`
}

// mobile モバイル用の写真URL
type mobile struct {
	L string `json:"l"`
}

// urls URL一覧
type urls struct {
	PC string `json:"pc"`
}
