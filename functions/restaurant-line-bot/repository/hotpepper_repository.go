package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"
)

type IHotpepperRepository interface {
	GetRestaurantInfos(response *model.HotpepperResponse, apiParams *model.APIParams) error
}

type hotpepperRepository struct {
}

func NewHotpepperRepository() IHotpepperRepository {
	return &hotpepperRepository{}
}

func (r *hotpepperRepository) GetRestaurantInfos(response *model.HotpepperResponse, apiParams *model.APIParams) error {
	url := utils.BuildAPIURL(apiParams)

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
