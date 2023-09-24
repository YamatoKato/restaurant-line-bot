package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
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

func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println(event, "event")
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

	signature := event.Headers["x-line-signature"]
	if signature == "" {
		signature = event.Headers["X-Line-Signature"]
	}
	if !validateSignature(os.Getenv("LINE_SECRET_TOKEN"), signature, []byte(event.Body)) {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       fmt.Sprintf(`{"message":"%s"}`+"\n", linebot.ErrInvalidSignature.Error()),
		}, nil
	}

	hotpepperRepository := repository.NewHotpepperRepository()
	usecase := usecase.NewUsecase(hotpepperRepository)
	controller := controller.NewController(usecase)

	if err := controller.HandleRequest(event, bot); err != nil {
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

func validateSignature(channelSecret string, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println(err, "validateSignature_base64.StdEncoding.DecodeString")
		return false
	}

	hash := hmac.New(sha256.New, []byte(channelSecret))
	_, err = hash.Write(body)
	if err != nil {
		fmt.Println(err, "validateSignature_hash.Write")
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}

func main() {
	lambda.Start(HandleRequest)
}
