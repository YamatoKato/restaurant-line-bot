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
	GENRE_LOOK_MORE_CODE         = ""
	GENRE_LOOK_MORE_JP           = "ジャンルは指定しない"
	GENRE_LOOK_MORE_IMG_URL      = "https://cdn.pixabay.com/photo/2017/06/01/18/46/cook-2364221_1280.jpg"
	GENRE_LOOK_MORE_DESC         = "ジャンル以外でも条件を指定して検索できます"
	ADD_CONDITION_WORDING        = "条件に追加する"
	PET_FRIENDLY_PARAM_VALUE     = "1" // ペット同伴可能
	PET_FRIENDLY_JP              = "ペット同伴可能"
	PET_FRIENDLY_IMG_URL         = "https://cdn.pixabay.com/photo/2023/03/24/16/41/people-7874368_1280.jpg"
	PET_FRIENDLY_DESC            = "ペットと同伴可能なお店を検索します"
	PET_FRIENDLY_PARAM_KEY       = "pet"
	TERRACE_PARAM_VALUE          = "1" // テラス席あり
	TERRACE_PARAM_KEY            = "terrace"
	TERRACE_JP                   = "テラス席あり"
	TERRACE_IMG_URL              = "https://cdn.pixabay.com/photo/2016/01/21/17/50/coffee-shop-1154289_1280.jpg"
	TERRACE_DESC                 = "テラス席があるお店を検索します"
	BUDGET_RANGE1_CODE           = "B011" // 予算1001～1500円
	BUDGET_RANGE2_CODE           = "B001" // 予算1501～2000円
	BUDGET_RANGE3_CODE           = "B002" // 予算2001～3000円
	BUDGET_RANGE4_CODE           = "B003" // 予算3001～4000円
	BUDGET_RANGE_KEY             = "budget"
	BUDGET_RANGE1_JP             = "1001～1500円"
	BUDGET_RANGE2_JP             = "1501～2000円"
	BUDGET_RANGE3_JP             = "2001～3000円"
	BUDGET_RANGE4_JP             = "3001～4000円"
	BUDGET_IMG_URL               = "https://cdn.pixabay.com/photo/2017/08/22/10/54/wallet-2668502_1280.jpg"
	BUDGET_JP                    = "予算"
	ADD_CONDITION_BUDGET_WORDING = "予算を指定する"
	BUDGET_DESC                  = "予算を指定して検索します"
	BUDGET_PARAM_KEY             = "budget"
	BUDGET_PARAM_VALUE           = "1"
	FREE_DRINK_PARAM_VALUE       = "1" // 飲み放題あり
	FREE_DRINK_PARAM_KEY         = "free_drink"
	FREE_DRINK_JP                = "飲み放題あり"
	FREE_DRINK_DESC              = "飲み放題があるお店を検索します"
	FREE_DRINK_IMG_URL           = "https://cdn.pixabay.com/photo/2016/06/03/21/44/folk-festival-1434523_1280.jpg"
	FREE_FOOD_PARAM_VALUE        = "1" // 食べ放題あり
	FREE_FOOD_PARAM_KEY          = "free_food"
	FREE_FOOD_JP                 = "食べ放題あり"
	FREE_FOOD_DESC               = "食べ放題があるお店を検索します"
	FREE_FOOD_IMG_URL            = "https://cdn.pixabay.com/photo/2017/03/23/15/48/buffet-2168675_1280.jpg"
	FREE_IMG_URL                 = "https://cdn.pixabay.com/photo/2016/08/03/10/19/mini-quiches-1566259_1280.jpg"
	FREE_DESC                    = "飲み放題・食べ放題またはその両方があるお店を検索します"
	PRIVATE_ROOM_PARAM_VALUE     = "1" // 個室あり
	PRIVATE_ROOM_PARAM_KEY       = "private_room"
	PRIVATE_ROOM_JP              = "個室あり"
	PRIVATE_ROOM_IMG_URL         = "https://cdn.pixabay.com/photo/2016/11/19/06/22/wine-1838132_1280.jpg"
	PRIVATE_ROOM_DESC            = "個室があるお店を検索します"
	MIDNIGHT_OPEN_PARAM_VALUE    = "1" // 23時以降も営業
	MIDNIGHT_OPEN_PARAM_KEY      = "midnight"
	MIDNIGHT_OPEN_JP             = "23時以降も営業"
	MIDNIGHT_OPEN_DESC           = "23時以降も営業しているお店を検索します"
	MIDNIGHT_MEAL_PARAM_VALUE    = "1"
	MIDNIGHT_MEAL_PARAM_KEY      = "midnight_meal"
	MIDNIGHT_MEAL_JP             = "23時以降食事OK"
	MIDNIGHT_MEAL_DESC           = "23時以降も食事ができるお店を検索します"
	MIDNIGHT_IMG_URL             = "https://cdn.pixabay.com/photo/2021/08/21/01/53/bars-6561626_1280.jpg"
	MIDNIGHT_DESC                = "23時以降も営業・食事またはその両方があるお店を検索します"
	PARKING_PARAM_VALUE          = "1" // 駐車場あり
	PARKING_PARAM_KEY            = "parking"
	PARKING_JP                   = "駐車場あり"
	PARKING_IMG_URL              = "https://cdn.pixabay.com/photo/2015/04/03/20/47/parking-705873_1280.jpg"
	PARKING_DESC                 = "駐車場があるお店を検索します"
	SMOKING_PARAM_VALUE          = "1" // 喫煙可
	SMOKING_PARAM_KEY            = "smoking"
	SMOKING_JP                   = "喫煙可"
	SMOKING_IMG_URL              = "https://cdn.pixabay.com/photo/2015/01/14/18/47/cigarette-599485_1280.jpg"
	SMOKING_DESC                 = "喫煙可のお店を検索します"
	INPUT_JP                     = "キーワード入力"
	INPUT_IMG_URL                = "https://cdn.pixabay.com/photo/2015/12/03/22/15/tablet-1075790_1280.jpg"
	INPUT_DESC                   = "キーワードを入力して検索します"
	INPUT_PARAM_KEY              = "keyword"
	ADD_CONDITION_INPUT_WORDING  = "入力する"
	INPUT_PARAM_VALUE            = "1"
	NON_CONDITION_WORDING        = "条件を指定しない"
	NON_CONDITION_IMG_URL        = "https://cdn.pixabay.com/photo/2017/10/14/09/56/journal-2850091_1280.jpg"
	NON_CONDITION_DESC           = "条件を指定せずに検索します"
)

// response APIレスポンス
type HotpepperResponse struct {
	Results Results `json:"results"`
}

// results APIレスポンスの内容
type Results struct {
	Shops []*Shop `json:"shop"`
}

// shop レストラン一覧
type Shop struct {
	Name       string   `json:"name"`
	Address    string   `json:"address"`
	Photo      photo    `json:"photo"`
	URLS       urls     `json:"urls"`
	NonSmoking string   `json:"non_smoking"`
	Budget     budget   `json:"budget"`
	Genre      category `json:"genre"`
	Access     string   `json:"access"`
	Open       string   `json:"open"`
	Close      string   `json:"close"`
	Catch      string   `json:"catch"`
}

type budget struct {
	Code    string `json:"code"`
	Name    string `json:"name"`
	Average string `json:"average"`
}

type category struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Catch string `json:"catch"`
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

type QueryParams struct {
	Lat       string
	Lng       string
	GenreCode string
	Keyword   string
	Address   string
}
