package controllers

import (
	"encoding/json"
	"os"
	"restaurant-line-bot/functions/restaurant-line-bot/domain/model"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/menudto"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/searchdto"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/interactor/usecase"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"
	"strconv"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sirupsen/logrus"
)

// LinebotController LINEBOTコントローラ
type LinebotController struct {
	searchInteractor  usecase.ISearchUsecase
	setMenuInteractor usecase.ISetMenuUsecase
	bot               *linebot.Client
}

// NewLinebotController コンストラクタ
func NewLinebotController(searchInteractor usecase.ISearchUsecase, setMenuInteractor usecase.ISetMenuUsecase) *LinebotController {
	bot, err := linebot.New(
		os.Getenv("LINE_SECRET_TOKEN"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		logrus.Fatalf("Error creating LINEBOT client: %v", err)
	}

	return &LinebotController{
		searchInteractor:  searchInteractor,
		setMenuInteractor: setMenuInteractor,
		bot:               bot,
	}
}

func (c *LinebotController) SetHelpMenu(e *linebot.Event) error {
	setHelpMenuInput := menudto.SetHelpMenuInput{
		ReplyToken: e.ReplyToken,
	}

	if _, err := c.setMenuInteractor.SetHelpMenu(setHelpMenuInput); err != nil {
		logrus.Errorf("Error setMenuInteractor.SetHelpMenu(setHelpMenuInput): %v", err)
		return err
	}
	return nil
}

func (c *LinebotController) SetAreaMenu(e *linebot.Event) error {
	setAreaMenuInput := menudto.SetAreaMenuInput{
		ReplyToken: e.ReplyToken,
	}

	if _, err := c.setMenuInteractor.SetAreaMenu(setAreaMenuInput); err != nil {
		logrus.Errorf("Error setMenuInteractor.SetAreaMenu(setAreaMenuInput): %v", err)
		return err
	}
	return nil
}

func (c *LinebotController) SetConfirmMenu(e *linebot.Event, messageType string, postbackStr string, areaMessage string) error {
	var data model.PostbackData

	if postbackStr == "" {
		// ユーザーがエリア名・位置情報を送信した場合
		areaWording := utils.GetAreaWording(areaMessage)
		data = model.PostbackData{
			AreaStr: areaWording,
		}

		if areaMessage == "" {
			// 位置情報を送信した場合
			msg := e.Message.(*linebot.LocationMessage)

			lat := strconv.FormatFloat(msg.Latitude, 'f', 2, 64)
			lng := strconv.FormatFloat(msg.Longitude, 'f', 2, 64)

			data = model.PostbackData{
				Lat:     lat,
				Lng:     lng,
				AreaStr: utils.GetAreaStrFromLocation(msg.Address),
			}
		}
	} else {
		// ユーザーがジャンル名・細かい条件を選択した場合
		if err := json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data); err != nil {
			logrus.Errorf("Error json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data): %v", err)
			return err
		}
	}

	setConfirmMenuInput := menudto.SetConfirmMenuInput{
		ReplyToken:   e.ReplyToken,
		AreaMessage:  areaMessage,
		PostbackData: data,
		MessageType:  messageType,
	}

	if _, err := c.setMenuInteractor.SetConfirmMenu(setConfirmMenuInput); err != nil {
		logrus.Errorf("Error setMenuInteractor.SetConfirmMenu(setConfirmMenuInput): %v", err)
		return err
	}

	return nil
}

func (c *LinebotController) SetGenreMenu(e *linebot.Event, postbackStr string) error {
	var data model.PostbackData
	if err := json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data); err != nil {
		logrus.Errorf("Error json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data): %v", err)
		return err
	}

	setGenreMenuInput := menudto.SetGenreMenuInput{
		ReplyToken:   e.ReplyToken,
		PostbackData: data,
	}

	if _, err := c.setMenuInteractor.SetGenreMenu(setGenreMenuInput); err != nil {
		logrus.Errorf("Error setMenuInteractor.SetGenreMenu(setGenreMenuInput): %v", err)
		return err
	}

	return nil
}

func (c *LinebotController) SetConditionMenu(e *linebot.Event, postbackStr string) error {
	var data model.PostbackData
	if err := json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data); err != nil {
		logrus.Errorf("Error json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data): %v", err)
		return err
	}

	setConditionMenuInput := menudto.SetConditionMenuInput{
		ReplyToken:   e.ReplyToken,
		PostbackData: data,
	}

	if _, err := c.setMenuInteractor.SetConditionMenu(setConditionMenuInput); err != nil {
		logrus.Errorf("Error setMenuInteractor.SetConditionMenu(setConditionMenuInput): %v", err)
		return err
	}

	return nil
}

func (c *LinebotController) GetRestaurantInfos(e *linebot.Event, postbackStr string) error {
	var data model.PostbackData
	if err := json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data); err != nil {
		logrus.Errorf("Error json.Unmarshal([]byte(utils.RemoveFirstTwoCharacters(postbackStr)), &data): %v", err)
		return err
	}

	getRestaurantInfosInput := searchdto.Input{
		ReplyToken:   e.ReplyToken,
		PostbackData: data,
	}

	if _, err := c.searchInteractor.GetRestaurantInfos(getRestaurantInfosInput); err != nil {
		logrus.Errorf("Error searchInteractor.GetRestaurantInfos(): %v", err)
		return err
	}

	return nil
}
