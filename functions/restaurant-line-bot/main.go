package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	echoadapter "github.com/awslabs/aws-lambda-go-api-proxy/echo"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var echoLambda *echoadapter.EchoLambda

func init() {
	e := echo.New()
	e.Use(middleware.Recover())

	r := InitDI(e)
	r.Init()

	echoLambda = echoadapter.New(e)
}

// API Gatewayのリクエストを echoLambda.ProxyWithContext を使用してEchoに転送し、EchoのレスポンスをLambdaからAPI Gatewayに返す
func HandleRequest(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return echoLambda.ProxyWithContext(ctx, event)

}

func main() {
	lambda.Start(HandleRequest)
}
