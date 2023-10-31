package interactor

import (
	"net/url"
	"restaurant-line-bot/functions/restaurant-line-bot/domain/model"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/searchdto"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/igateway"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/ipresenter"

	"github.com/sirupsen/logrus"
)

// SearchInteractor 検索インタラクタ
type SearchInteractor struct {
	hotpepperGateway igateway.IHotpepperGateway
	linePresenter    ipresenter.ILinePresenter
}

// NewSearchInteractor コンストラクタ
func NewSearchInteractor(
	hotpepperGateway igateway.IHotpepperGateway,
	linePresenter ipresenter.ILinePresenter) *SearchInteractor {
	return &SearchInteractor{
		hotpepperGateway: hotpepperGateway,
		linePresenter:    linePresenter,
	}
}

// GetRestaurantInfos 検索処理
func (interactor *SearchInteractor) GetRestaurantInfos(in searchdto.Input) (searchdto.Output, error) {
	queryParams := convertToQueryParams(in.PostbackData)
	res, err := interactor.hotpepperGateway.GetRestaurantInfos(queryParams)
	if err != nil {
		return searchdto.Output{}, err
	}

	searchdtoOutput := searchdto.Output{
		ReplyToken: in.ReplyToken,
		Response:   res,
		URIActionDataForMore: &model.URIActionData{
			Label: model.URI_LABEL_GET_RESTAURANT_INFOS_MORE,
		},
		URIActionDataForMap: &model.URIActionData{
			Label: model.URI_LABEL_GET_RESTAURANT_INFOS_MAP,
			URI:   model.GOOGLE_MAPS_URI,
		},
		// URIActionDataForTel: &model.URIActionData{
		// 	Label: model.URI_LABEL_GET_RESTAURANT_INFOS_TEL,
		// },
		TemplateMessageData: &model.TemplateMessageData{
			AltText: model.TM_ALT_TEXT_GET_RESTAURANT_INFOS,
		},
		TextMessageData: &model.TextMessageData{
			Content: model.TM_CONTENT_GET_RESTAURANT_INFOS_NOTHING_HIT,
		},
	}

	// 空の場合
	if len(res.Results.Shops) == 0 {
		logrus.Info("店舗が見つかりませんでした。")
		if err := interactor.linePresenter.SetNothingHitReplyMessage(searchdtoOutput); err != nil {
			return searchdto.Output{}, err
		}
	}

	if err := interactor.linePresenter.SetRestaurantListReplyMessage(searchdtoOutput); err != nil {
		return searchdto.Output{}, err
	}

	return searchdto.Output{}, nil
}

func convertToQueryParams(data model.PostbackData) string {
	values := url.Values{}
	if data.AreaStr != "" {
		values.Add("keyword", data.AreaStr)
	}

	if data.Keyword != "" && values.Get("keyword") == "" {
		values.Add("keyword", data.Keyword)
	} else if data.Keyword != "" && values.Get("keyword") != "" {
		values.Set("keyword", values.Get("keyword")+" "+data.Keyword)
	}

	if data.Smoking != "" && values.Get("keyword") == "" {
		values.Add("keyword", "喫煙")
	} else if data.Smoking != "" && values.Get("keyword") != "" {
		values.Set("keyword", values.Get("keyword")+" 喫煙")
	}

	if data.Lat != "" {
		values.Add("lat", data.Lat)
	}
	if data.Lng != "" {
		values.Add("lng", data.Lng)
	}
	if data.GenreCode != "" {
		values.Add("genre", data.GenreCode)
	}
	if data.Parking != "" {
		values.Add("parking", data.Parking)
	}
	if data.PetFriendly != "" {
		values.Add("pet", data.PetFriendly)
	}
	if data.MidnightOpen != "" {
		values.Add("midnight", data.MidnightOpen)
	}
	if data.MidnightMeal != "" {
		values.Add("midnight_meal", data.MidnightMeal)
	}
	if data.PrivateRoom != "" {
		values.Add("private_room", data.PrivateRoom)
	}
	if data.FreeFood != "" {
		values.Add("free_food", data.FreeFood)
	}
	if data.FreeDrink != "" {
		values.Add("free_drink", data.FreeDrink)
	}
	if data.Budget != "" {
		values.Add("budget", data.Budget)
	}
	if data.Terrace != "" {
		values.Add("terrace", data.Terrace)
	}

	return values.Encode()
}
