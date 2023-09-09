package usecase

import (
	"fmt"
	"net/url"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/repository"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type IBotUsecase interface {
	GetRestaurantInfos(area model.Area) (*linebot.TemplateMessage, error)
}

type botUsecase struct {
	hr repository.IHotpepperRepository
}

func NewBotUsecase(hr repository.IHotpepperRepository) IBotUsecase {
	return &botUsecase{hr}
}

func (bu *botUsecase) GetRestaurantInfos(area model.Area) (*linebot.TemplateMessage, error) {
	response := model.HotpepperResponse{}
	var ccs []*linebot.CarouselColumn

	if err := bu.hr.GetRestaurantInfos(&response, &area); err != nil {
		fmt.Println(err, "bot_usecase@GetRestaurantInfos-bu.hr.GetRestaurantInfos")
		return nil, err
	}
	for _, shop := range response.Results.Shop {
		addr := shop.Address
		// 61文字以上ある場合はそれ以降をカット
		addr = utils.CutString(addr, 61)

		cc := linebot.NewCarouselColumn(
			shop.Photo.Mobile.L,
			shop.Name,
			addr,
			linebot.NewURIAction("ホットペッパーで開く", shop.URLS.PC),
			linebot.NewURIAction("GoogleMapで開く", "https://www.google.com/maps/search/?api=1&query="+url.QueryEscape(utils.RemoveSpaces(shop.Name)+" "+utils.RemoveSpaces(shop.Address))),
		).WithImageOptions("#FFFFFF")
		ccs = append(ccs, cc)
	}

	res := linebot.NewTemplateMessage(
		"レストラン一覧",
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)
	return res, nil

}
