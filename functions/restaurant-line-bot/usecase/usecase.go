package usecase

import (
	"fmt"
	"net/url"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/repository"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type IUsecase interface {
	GetRestaurantInfos(area model.Area) (*linebot.TemplateMessage, error)
	SetAreaMenu() (*linebot.TemplateMessage, error)
}

type usecase struct {
	hr repository.IHotpepperRepository
}

func NewUsecase(hr repository.IHotpepperRepository) IUsecase {
	return &usecase{hr}
}

func (u *usecase) GetRestaurantInfos(area model.Area) (*linebot.TemplateMessage, error) {
	response := model.HotpepperResponse{}
	var ccs []*linebot.CarouselColumn

	if err := u.hr.GetRestaurantInfos(&response, &area); err != nil {
		fmt.Println(err, "usecase@GetRestaurantInfos-bu.hr.GetRestaurantInfos")
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

func (u *usecase) SetAreaMenu() (*linebot.TemplateMessage, error) {
	bt := linebot.NewButtonsTemplate(
		"https://cdn.pixabay.com/photo/2017/08/17/07/47/travel-2650303_1280.jpg",
		"１/２　検索エリアを指定",
		"指定エリアまたは位置情報を送信し、その中心から検索します\n※エリアを指定する場合「---」を消さずに入力してください",
		linebot.NewPostbackAction("エリアを指定する", "input_area", "", "", "openKeyboard", "---\n都道府県: 東京\n地区: 渋谷\n---"),
		linebot.NewURIAction("位置情報を送る", "https://line.me/R/nv/location/"),
	)
	tm := linebot.NewTemplateMessage("エリアを選択してください", bt)
	return tm, nil
}
