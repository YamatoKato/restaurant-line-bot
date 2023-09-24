package model

type Genre struct {
	Code   string
	Name   string
	ImgURL string
}

func CreateGenre(code, name, imgURL string) Genre {
	return Genre{
		Code:   code,
		Name:   name,
		ImgURL: imgURL,
	}
}

func SearchGenreNameByCode(genreCode string) string {
	switch genreCode {
	case GENRE_JAPANESE_CODE:
		return GENRE_JAPANESE_JP
	case GENRE_WESTERN_CODE:
		return GENRE_WESTERN_JP
	case GENRE_CHINESE_CODE:
		return GENRE_CHINESE_JP
	case GENRE_ITALIAN_FRENCH_CODE:
		return GENRE_ITALIAN_FRENCH_JP
	case GENRE_KOREAN_CODE:
		return GENRE_KOREAN_JP
	case GENRE_RAMEN_CODE:
		return GENRE_RAMEN_JP
	case GENRE_YAKINIKU_OFFAL_CODE:
		return GENRE_YAKINIKU_OFFAL_JP
	case GENRE_CAFE_SWEETS_CODE:
		return GENRE_CAFE_SWEETS_JP
	case GENRE_IZAKAYA_CODE:
		return GENRE_IZAKAYA_JP
	case GENRE_ROULETTE_CODE:
		return GENRE_ROULETTE_JP
	case GENRE_LOOK_MORE_CODE:
		return GENRE_LOOK_MORE_JP
	// 他のジャンルに対する定義を追加
	default:
		return ""
	}
}

func SearchGenreImgUrlByCode(genreCode string) string {
	switch genreCode {
	case GENRE_JAPANESE_CODE:
		return GENRE_JAPANESE_IMG_URL
	case GENRE_WESTERN_CODE:
		return GENRE_WESTERN_IMG_URL
	case GENRE_CHINESE_CODE:
		return GENRE_CHINESE_IMG_URL
	case GENRE_ITALIAN_FRENCH_CODE:
		return GENRE_ITALIAN_FRENCH_IMG_URL
	case GENRE_KOREAN_CODE:
		return GENRE_KOREAN_IMG_URL
	case GENRE_RAMEN_CODE:
		return GENRE_RAMEN_IMG_URL
	case GENRE_YAKINIKU_OFFAL_CODE:
		return GENRE_YAKINIKU_OFFAL_IMG_URL
	case GENRE_CAFE_SWEETS_CODE:
		return GENRE_CAFE_SWEETS_IMG_URL
	case GENRE_IZAKAYA_CODE:
		return GENRE_IZAKAYA_IMG_URL
	// 他のジャンルに対する定義を追加
	case GENRE_ROULETTE_CODE:
		return GENRE_ROULETTE_IMG_URL
	case GENRE_LOOK_MORE_CODE:
		return GENRE_LOOK_MORE_IMG_URL
	default:
		return ""
	}
}
