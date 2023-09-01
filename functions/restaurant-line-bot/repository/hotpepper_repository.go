package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
)

const (
	API_URL_BY_AREA = "https://webservice.recruit.co.jp/hotpepper/gourmet/v1/?format=json&key=%s&lat=%s&lng=%s"
)

type IHotpepperRepository interface {
	GetRestaurantInfos(response *model.HotpepperResponse, area *model.Area) error
}

type hotpepperRepository struct {
}

func NewHotpepperRepository() IHotpepperRepository {
	return &hotpepperRepository{}
}

func (r *hotpepperRepository) GetRestaurantInfos(response *model.HotpepperResponse, area *model.Area) error {
	url := fmt.Sprintf(API_URL_BY_AREA, os.Getenv("HOTPEPPER_API_KEY"), area.Latitude, area.Longitude)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err, "hotpepper_repository@GetRestaurantInfos_http.Get")
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err, "hotpepper_repository@GetRestaurantInfos_ioutil.ReadAll")
		return err
	}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println(err, "hotpepper_repository@GetRestaurantInfos_json.Unmarshal")
		return err
	}

	defer resp.Body.Close()
	return nil
}
