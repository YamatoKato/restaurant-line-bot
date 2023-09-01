package usecase

import (
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/repository"
	"unicode/utf8"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type IBotUsecase interface {
	GetRestaurantInfos(area model.Area) ([]*linebot.CarouselColumn, error)
}

type botUsecase struct {
	hr repository.IHotpepperRepository
}

func NewBotUsecase(hr repository.IHotpepperRepository) IBotUsecase {
	return &botUsecase{hr}
}

func (bu *botUsecase) GetRestaurantInfos(area model.Area) ([]*linebot.CarouselColumn, error) {
	response := model.HotpepperResponse{}
	var ccs []*linebot.CarouselColumn

	if err := bu.hr.GetRestaurantInfos(&response, &area); err != nil {
		fmt.Println(err, "bot_usecase@GetRestaurantInfos-bu.hr.GetRestaurantInfos")
		return nil, err
	}
	for _, shop := range response.Results.Shop {
		addr := shop.Address
		// 61文字以上ある場合はそれ以降をカット
		if 60 < utf8.RuneCountInString(addr) {
			addr = string([]rune(addr)[:60])
		}

		cc := linebot.NewCarouselColumn(
			shop.Photo.Mobile.L,
			shop.Name,
			addr,
			linebot.NewURIAction("ホットペッパーで開く", shop.URLS.PC),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}
	return ccs, nil

}
