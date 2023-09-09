# アプリケーション概要

このアプリケーションは、LINE Messaging API を活用したレストラン検索チャットボットです。ユーザーは LINE アプリを通じて、レストラン情報を検索し、詳細を表示することができます。

## 主要コンポーネント

### LINE Messaging API

LINE Messaging API は、LINE プラットフォーム上でユーザーとのコミュニケーションを実現するための API です。このアプリケーションでは、LINE ユーザーからのメッセージを受信し、応答メッセージを生成するために使用されています。

### AWS SAM (Serverless Application Model)

AWS SAM を使用して APIGateway,Lambda 関数のデプロイと設定を管理しています。

### HotpepperAPI

レストラン情報を提供する外部 API です。このアプリケーションでは、ユーザーからの検索要求に応じて HotpepperAPI にクエリを送信し、レストラン情報を取得します。

## 動作フロー

1. ユーザーは LINE アプリを通じて、チャットボットにメッセージを送信します。
2. LINE Messaging API はユーザーのメッセージを受信し、AWS API Gateway を介して Lambda 関数に転送します。
3. Lambda 関数はユーザーの検索要求を処理し、HotpepperAPI に対するクエリを生成します。
4. HotpepperAPI からのレストラン情報を受け取り、ユーザーに返信メッセージを生成します。
5. 返信メッセージは LINE Messaging API を介してユーザーに送信され、チャットボットの応答が表示されます。

このアプリケーションは、LINE プラットフォームと AWS サーバーレステクノロジーを組み合わせて、レストラン検索の簡単で迅速な方法を提供します。

# アーキテクチャ

main.go->controller.go->usecase.go->repository.go

※エンドポイントが単一のため、router は不要と判断。

## main.go

- lambda の`/restaurant＠POST`リクエストを受け取る。
- line-bot-sdk の初期化
- DI
- 署名の検証
- `@return` リクエストのレスポンス内容を返す

## controller.go

- リクエストから Webhook イベントを取得
- Webhook イベントを解析し、適切な usecase.go をトリガー
- line-bot-sdk によって LINEMessagingAPI の各種エンドポイントをトリガー
- `@return` 上記処理のエラー

## usecase.go

- repository.go にて取得したデータを line-bot-sdk の各種構造体に加工。
- `@return` line-bot-sdk の各種構造体

## repository.go

- 今回は、HotpepperAPI を永続化されたデータとして扱い、API を利用してデータを取得する
- `@return` 取得したデータを返す

## model.go

- 各階層で利用されるデータ構造モデルを定義
