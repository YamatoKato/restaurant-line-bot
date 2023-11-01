package infrastructure

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"restaurant-line-bot/functions/restaurant-line-bot/domain/model"
	"restaurant-line-bot/functions/restaurant-line-bot/interfaces/controllers"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/aws/aws-lambda-go/events"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sirupsen/logrus"
)

// Router ルーティング
type Router struct {
	lc *controllers.LinebotController
}

// NewRouter コンストラクタ
func NewRouter(lc *controllers.LinebotController) *Router {
	return &Router{lc: lc}
}

// イベントごとにルーティング
func (r *Router) CatchEvents(event events.APIGatewayProxyRequest) error {
	// 署名の検証
	signature := event.Headers["x-line-signature"]
	if signature == "" {
		signature = event.Headers["X-Line-Signature"]
	}
	if !validateSignature(os.Getenv("LINE_SECRET_TOKEN"), signature, []byte(event.Body)) {
		logrus.Error("署名の検証に失敗しました")
		return fmt.Errorf("署名の検証に失敗しました")
	}

	webhook := model.Webhook{}

	// リクエストからイベントを取得
	if err := json.Unmarshal([]byte(event.Body), &webhook); err != nil {
		logrus.Error(err, "router@CatchEvents_json.Unmarshal")
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
					// 最初の導入メッセージ
					if err := r.lc.SetAreaMenu(we); err != nil {
						logrus.Error(err, "router@SetAreaMenu")
						return err
					}
					return nil
				} else if utils.ContainsHyphen(userMessage) {
					// メッセージにハイフンを含む場合（エリア指定）
					if err := r.lc.SetConfirmMenu(we, model.PBD_PREFIX_IDENTIFY_GENRE, "", userMessage); err != nil {
						logrus.Error(err, "router@SetConfirmMenu")
						return err
					}
					return nil
				} else {
					// それ以外の場合（エリア未指定）
					if err := r.lc.SetHelpMenu(we); err != nil {
						logrus.Error(err, "router@SetHelpMenu")
						return err
					}
					return nil
				}

			// メッセージが位置情報の場合
			case *linebot.LocationMessage:
				if err := r.lc.SetConfirmMenu(we, model.PBD_PREFIX_IDENTIFY_GENRE, "", ""); err != nil {
					logrus.Error(err, "router@SetConfirmMenu")
					return err
				}
				return nil
			}
		} else if we.Type == linebot.EventTypePostback {

			// 確認ボタン
			if model.PBD_PREFIX_IDENTIFY_CONFIRM == model.GetPrefix(we.Postback.Data) {
				if err := r.lc.SetConfirmMenu(we, "", we.Postback.Data, ""); err != nil {
					logrus.Error(err, "router@SetConfirmMenu")
					return err
				}
			}

			// ジャンルボタン
			if model.PBD_PREFIX_IDENTIFY_GENRE == model.GetPrefix(we.Postback.Data) {
				if err := r.lc.SetGenreMenu(we, we.Postback.Data); err != nil {
					logrus.Error(err, "router@SetGenreMenu")
					return err
				}
			}

			// 条件ボタン
			if model.PBD_PREFIX_IDENTIFY_CONDITION == model.GetPrefix(we.Postback.Data) {
				if err := r.lc.SetConditionMenu(we, we.Postback.Data); err != nil {
					logrus.Error(err, "router@SetConditionMenu")
					return err
				}
			}

			// お店一覧
			if model.PBD_PREFIX_IDENTIFY_SEARCH == model.GetPrefix(we.Postback.Data) {
				if err := r.lc.GetRestaurantInfos(we, we.Postback.Data); err != nil {
					logrus.Error(err, "router@GetRestaurantInfos")
					return err
				}
			}

			return nil

		} else {
			if err := r.lc.SetHelpMenu(we); err != nil {
				logrus.Error(err, "router@SetHelpMenu")
				return err
			}
			return nil
		}
	}

	return nil

}

func validateSignature(channelSecret string, signature string, body []byte) bool {
	decoded, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		logrus.Error(err, "validateSignature_base64.StdEncoding.DecodeString")
		return false
	}

	hash := hmac.New(sha256.New, []byte(channelSecret))
	_, err = hash.Write(body)
	if err != nil {
		logrus.Error(err, "validateSignature_hash.Write")
		return false
	}

	return hmac.Equal(decoded, hash.Sum(nil))
}
