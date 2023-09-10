package usecase

import (
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (u *usecase) SetGenreMenu(apiParams *model.APIParams) (*linebot.TemplateMessage, error) {
	var ccs []*linebot.CarouselColumn
	genres := []model.Genre{
		{
			Code:   model.GENRE_JAPANESE_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_JAPANESE_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_JAPANESE_CODE),
		},
		{
			Code:   model.GENRE_WESTERN_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_WESTERN_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_WESTERN_CODE),
		},
		{
			Code:   model.GENRE_CHINESE_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_CHINESE_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_CHINESE_CODE),
		},
		{
			Code:   model.GENRE_ITALIAN_FRENCH_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_ITALIAN_FRENCH_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_ITALIAN_FRENCH_CODE),
		},
		{
			Code:   model.GENRE_KOREAN_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_KOREAN_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_KOREAN_CODE),
		},
		{
			Code:   model.GENRE_RAMEN_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_RAMEN_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_RAMEN_CODE),
		},
		{
			Code:   model.GENRE_YAKINIKU_OFFAL_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_YAKINIKU_OFFAL_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_YAKINIKU_OFFAL_CODE),
		},
		{
			Code:   model.GENRE_CAFE_SWEETS_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_CAFE_SWEETS_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_CAFE_SWEETS_CODE),
		},
		{
			Code:   model.GENRE_IZAKAYA_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_IZAKAYA_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_IZAKAYA_CODE),
		},
		{
			Code:   model.GENRE_ROULETTE_CODE,
			Name:   utils.SearchGenreNameByCode(model.GENRE_ROULETTE_CODE),
			ImgURL: utils.SearchGenreImgUrlByCode(model.GENRE_ROULETTE_CODE),
		},
	}

	for _, genre := range genres {
		postBackData := utils.CreatePostBackData(apiParams, genre.Code)
		cc := linebot.NewCarouselColumn(
			genre.ImgURL,
			genre.Name,
			fmt.Sprintf("エリア「%s」付近", apiParams.AreaStr),
			linebot.NewPostbackAction(fmt.Sprintf("ジャンル「%s」で検索", genre.Name), postBackData, "", "", "", ""),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}

	tm := linebot.NewTemplateMessage(
		"ジャンル選択",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)
	return tm, nil
}
