package gateway

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"restaurant-line-bot/functions/restaurant-line-bot/domain/model"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/igateway"

	"github.com/sirupsen/logrus"
)

const BASE_HOTPEPPER_API_URL = "https://webservice.recruit.co.jp/hotpepper/gourmet/v1/"

type HotpepperGateway struct {
	baseApiURL string
}

func NewHotpepperGateway() igateway.IHotpepperGateway {
	params := url.Values{}
	params.Add("format", "json")
	params.Add("key", os.Getenv("HOTPEPPER_API_KEY"))
	baseApiURL := BASE_HOTPEPPER_API_URL + "?" + params.Encode()

	return &HotpepperGateway{
		baseApiURL: baseApiURL,
	}
}

func (g *HotpepperGateway) GetRestaurantInfos(apiParams string) (*model.HotpepperResponse, error) {
	url := g.baseApiURL + "&" + apiParams
	var response *model.HotpepperResponse

	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		logrus.Errorf("Error getting response: %v", err)
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("Error reading response body: %v", err)
		return nil, err
	}

	// jsonでログを表示
	logrus.Infof("Response body: %v", string(body))

	if err := json.Unmarshal(body, &response); err != nil {
		logrus.Errorf("Error unmarshaling response body: %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	return response, nil
}
