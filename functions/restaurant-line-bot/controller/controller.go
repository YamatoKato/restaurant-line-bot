package controller

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/usecase"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/aws/aws-lambda-go/events"
)

type IController interface {
	HandleRequest(event events.APIGatewayProxyRequest, bot *linebot.Client) error
}

type Controller struct {
	u usecase.IUsecase
}

func NewController(u usecase.IUsecase) IController {
	return &Controller{u}
}

func (c *Controller) HandleRequest(event events.APIGatewayProxyRequest, bot *linebot.Client) error {
	webhook := model.Webhook{}

	// リクエストからイベントを取得
	if err := json.Unmarshal([]byte(event.Body), &webhook); err != nil {
		fmt.Println(err, "controller@HandleRequest_json.Unmarshal")
		return err
	}

	for _, we := range webhook.Events {

		// イベントがメッセージの受信だった場合
		if we.Type == linebot.EventTypeMessage {

			switch message := we.Message.(type) {

			// メッセージがテキスト形式の場合
			case *linebot.TextMessage:
				userMessage := message.Text

				if userMessage == model.INTRO_MESSAGE {
					if err := setAreaMenu(c, we, bot); err != nil {
						fmt.Println(err, "controller@setAreaMenu")
						return err
					}
					return nil
				} else if utils.ContainsHyphen(userMessage) {
					// エリア指定の場合
					if err := setGenreMenu(c, we, bot, userMessage); err != nil {
						fmt.Println(err, "controller@setGenreMenu")
						return err
					}
					return nil
				}
			// メッセージが位置情報の場合
			case *linebot.LocationMessage:
				if err := setGenreMenu(c, we, bot, ""); err != nil {
					fmt.Println(err, "controller@*linebot.LocationMessage_setGenreMenu")
					return err
				}
				// if err := getRestaurantInfos(c, we, bot); err != nil {
				// 	fmt.Println(err, "controller@getRestaurantInfos")
				// 	return err
				// }
				return nil
			}
		}
	}

	return nil
}
