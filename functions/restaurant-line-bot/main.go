package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/line/line-bot-sdk-go/v7/linebot"

	"restaurant-line-bot/functions/restaurant-line-bot/controller"
	"restaurant-line-bot/functions/restaurant-line-bot/repository"
	"restaurant-line-bot/functions/restaurant-line-bot/usecase"
)

type Webhook struct {
	Destination string           `json:"destination"`
	Events      []*linebot.Event `json:"events"`
}

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println(event.Body, "event.Body")
	// BOTを初期化
	bot, err := linebot.New(
		os.Getenv("LINE_SECRET_TOKEN"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		fmt.Println(err, "linebot.New")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"message":"%s"}`+"\n", http.StatusText(http.StatusInternalServerError)),
		}, nil
	}

	hotpepperRepository := repository.NewHotpepperRepository()
	botUsecase := usecase.NewBotUsecase(hotpepperRepository)
	botController := controller.NewController(botUsecase)

	if err := botController.HandleRequest(event, bot); err != nil {
		fmt.Println(err, "main.botController.HandleRequest")
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       fmt.Sprintf(`{"message":"%s"}`+"\n", http.StatusText(http.StatusInternalServerError)),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
