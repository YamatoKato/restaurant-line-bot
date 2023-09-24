package usecase

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func (u *usecase) SetGenreMenu(postbackData *model.PostbackData) ([]linebot.SendingMessage, error) {
	var ccs []*linebot.CarouselColumn
	fmt.Println(postbackData.AreaStr)
	genres := []model.Genre{
		model.CreateGenre(model.GENRE_JAPANESE_CODE, model.GENRE_JAPANESE_JP, model.GENRE_JAPANESE_IMG_URL),
		model.CreateGenre(model.GENRE_WESTERN_CODE, model.GENRE_WESTERN_JP, model.GENRE_WESTERN_IMG_URL),
		model.CreateGenre(model.GENRE_CHINESE_CODE, model.GENRE_CHINESE_JP, model.GENRE_CHINESE_IMG_URL),
		model.CreateGenre(model.GENRE_ITALIAN_FRENCH_CODE, model.GENRE_ITALIAN_FRENCH_JP, model.GENRE_ITALIAN_FRENCH_IMG_URL),
		model.CreateGenre(model.GENRE_KOREAN_CODE, model.GENRE_KOREAN_JP, model.GENRE_KOREAN_IMG_URL),
		model.CreateGenre(model.GENRE_RAMEN_CODE, model.GENRE_RAMEN_JP, model.GENRE_RAMEN_IMG_URL),
		model.CreateGenre(model.GENRE_YAKINIKU_OFFAL_CODE, model.GENRE_YAKINIKU_OFFAL_JP, model.GENRE_YAKINIKU_OFFAL_IMG_URL),
		model.CreateGenre(model.GENRE_CAFE_SWEETS_CODE, model.GENRE_CAFE_SWEETS_JP, model.GENRE_CAFE_SWEETS_IMG_URL),
		model.CreateGenre(model.GENRE_IZAKAYA_CODE, model.GENRE_IZAKAYA_JP, model.GENRE_IZAKAYA_IMG_URL),
		model.CreateGenre(model.GENRE_LOOK_MORE_CODE, model.GENRE_LOOK_MORE_JP, model.GENRE_LOOK_MORE_IMG_URL),
	}

	for i, genre := range genres {
		postbackData.GenreCode = genre.Code
		jsonData, err := json.Marshal(postbackData)
		if err != nil {
			fmt.Println(err, "usecase.set_genre_menu.json.Marshal(postbackData)")
			return nil, err
		}

		// ジャンルを指定しない
		if i == len(genres)-1 {
			cc := linebot.NewCarouselColumn(
				genre.ImgURL,
				genre.Name,
				utils.CutString(fmt.Sprintf("ジャンル「なし」\nエリア「%s」付近", postbackData.AreaStr), 61),
				linebot.NewPostbackAction(genre.Name, model.PBD_PREFIX_IDENTIFY_CONFIRM+" "+string(jsonData), "", "ジャンルを指定しない\n\n※ジャンル選択を変更したい場合は、↑のジャンル一覧から再度選択してください", "", ""),
			).WithImageOptions("#FFFFFF")
			ccs = append(ccs, cc)
		} else {
			// ジャンル指定する
			cc := linebot.NewCarouselColumn(
				genre.ImgURL,
				genre.Name,
				utils.CutString(fmt.Sprintf("ジャンル「%s」\nエリア「%s」付近", genre.Name, postbackData.AreaStr), 61),
				linebot.NewPostbackAction(genre.Name, model.PBD_PREFIX_IDENTIFY_CONFIRM+" "+string(jsonData), "", fmt.Sprintf("ジャンル「%s」で選択しました\n\n※ジャンル選択を変更したい場合は、↑のジャンル一覧から再度選択してください", genre.Name), "", ""),
			).WithImageOptions("#FFFFFF")
			ccs = append(ccs, cc)
		}
	}

	var sms []linebot.SendingMessage

	tm := linebot.NewTemplateMessage(
		"ジャンル選択",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)

	sms = append(sms, tm)
	return sms, nil
}
