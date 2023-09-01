package controller

import (
	"encoding/json"
	"fmt"
	"restaurant-line-bot/functions/restaurant-line-bot/model"
	"restaurant-line-bot/functions/restaurant-line-bot/usecase"

	"github.com/line/line-bot-sdk-go/v7/linebot"

	"github.com/aws/aws-lambda-go/events"
)

type IController interface {
	HandleRequest(event events.APIGatewayProxyRequest, bot *linebot.Client) error
}

type Controller struct {
	bu usecase.IBotUsecase
}

func NewController(bu usecase.IBotUsecase) IController {
	return &Controller{bu}
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
				replyMessage := message.Text
				if _, err := bot.ReplyMessage(we.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
					fmt.Println(err, "controller@HandleRequest_bot.ReplyMessage")
					return err
				}
				// メッセージが位置情報の場合

			case *linebot.LocationMessage:
				if err := GetRestaurantInfos(c, we, bot); err != nil {
					fmt.Println(err, "controller@getRestaurantInfos")
					return err
				}
			}
		}
	}

	return nil
}
