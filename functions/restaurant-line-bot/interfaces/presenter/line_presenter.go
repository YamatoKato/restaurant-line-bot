package presenter

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"restaurant-line-bot/functions/restaurant-line-bot/domain/model"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/menudto"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/dto/searchdto"
	"restaurant-line-bot/functions/restaurant-line-bot/usecases/ipresenter"
	"restaurant-line-bot/functions/restaurant-line-bot/utils"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/sirupsen/logrus"
)

// LinePresenter LINEプレゼンタ
type LinePresenter struct {
	bot *linebot.Client
}

// NewLinePresenter コンストラクタ
func NewLinePresenter() ipresenter.ILinePresenter {
	bot, err := linebot.New(
		os.Getenv("LINE_SECRET_TOKEN"),
		os.Getenv("LINE_ACCESS_TOKEN"),
	)
	if err != nil {
		logrus.Fatalf("Error creating LINEBOT client: %v", err)
	}

	return &LinePresenter{bot: bot}
}

// SetHelpMenuReplyMessage リプライメッセージ
func (p *LinePresenter) SetHelpMenuReplyMessage(out menudto.SetHelpMenuOutput) error {
	if _, err := p.bot.ReplyMessage(out.ReplyToken, linebot.NewTextMessage(out.TextMessageData.Content)).Do(); err != nil {
		logrus.Errorf("Error LINEBOT replying message: %v", err)
		return err
	}
	return nil
}

// SetAreaMenu リプライメッセージ
func (p *LinePresenter) SetAreaMenuReplyMessage(out menudto.SetAreaMenuOutput) error {
	bt := linebot.NewButtonsTemplate(
		out.ButtonsTemplateData.ThumbnailImageURL,
		out.ButtonsTemplateData.Title,
		out.ButtonsTemplateData.Text,
		linebot.NewPostbackAction(
			out.PostbackActionData.Label,
			out.PostbackActionData.Data,
			out.PostbackActionData.Text,
			out.PostbackActionData.DisplayText,
			out.PostbackActionData.InputOption,
			out.PostbackActionData.FillInText,
		),
		linebot.NewURIAction(
			out.URIActionData.Label,
			out.URIActionData.URI,
		),
	)
	tm := linebot.NewTemplateMessage(out.TemplateMessageData.AltText, bt)

	if _, err := p.bot.ReplyMessage(out.ReplyToken, tm).Do(); err != nil {
		logrus.Errorf("Error LINEBOT replying message: %v", err)
		return err
	}
	return nil
}

// SetConfirmMenu リプライメッセージ
func (p *LinePresenter) SetConfirmMenuReplyMessage(out menudto.SetConfirmMenuOutput) error {
	jsonData, err := json.Marshal(out.PostbackData)
	if err != nil {
		logrus.Errorf("Error json.Marshal(out.PostbackData): %v", err)
		return err
	}

	bt := linebot.NewButtonsTemplate(
		out.ButtonsTemplateData.ThumbnailImageURL,
		out.ButtonsTemplateData.Title,
		out.ButtonsTemplateData.Text,
		linebot.NewPostbackAction(
			out.PostbackActionData.Label,
			out.PostbackActionData.Data+string(jsonData),
			out.PostbackActionData.Text,
			out.PostbackActionData.DisplayText,
			out.PostbackActionData.InputOption,
			out.PostbackActionData.FillInText,
		),
		linebot.NewPostbackAction(
			out.PostbackActionDataForSearch.Label,
			out.PostbackActionDataForSearch.Data+string(jsonData),
			out.PostbackActionDataForSearch.Text,
			out.PostbackActionDataForSearch.DisplayText,
			out.PostbackActionDataForSearch.InputOption,
			out.PostbackActionDataForSearch.FillInText,
		),
	)
	tm := linebot.NewTemplateMessage(out.TemplateMessageData.AltText, bt)

	if _, err := p.bot.ReplyMessage(out.ReplyToken, tm).Do(); err != nil {
		logrus.Errorf("Error LINEBOT replying message: %v", err)
		return err
	}
	return nil
}

func (p *LinePresenter) SetGenreMenuReplyMessage(out menudto.SetGenreMenuOutput) error {
	var ccs []*linebot.CarouselColumn
	for i, genre := range out.Genres {
		newPostbackData := model.SetPostbackDataField(out.PostbackData, "genreCode", genre.Code)
		jsonData, err := json.Marshal(newPostbackData)
		if err != nil {
			logrus.Errorf("Error json.Marshal(newPostbackData): %v", err)
			return err
		}
		if i == len(out.Genres)-1 {
			// ジャンルを指定しない
			cc := linebot.NewCarouselColumn(
				genre.ImgURL,
				genre.Name,
				utils.CutString(fmt.Sprintf(out.CarouselColumnData.Text, genre.Name, out.AreaStr), 61),
				linebot.NewPostbackAction(
					genre.Name,
					out.PostbackActionData.Data+string(jsonData),
					out.PostbackActionData.Text,
					fmt.Sprintf(out.PostbackActionData.DisplayText, genre.Name),
					"",
					out.PostbackActionData.FillInText,
				),
			).WithImageOptions("#FFFFFF")
			ccs = append(ccs, cc)

		} else {
			// ジャンル指定する
			cc := linebot.NewCarouselColumn(
				genre.ImgURL,
				genre.Name,
				utils.CutString(fmt.Sprintf(out.CarouselColumnData.Text, genre.Name, out.AreaStr), 61),
				linebot.NewPostbackAction(
					genre.Name,
					out.PostbackActionData.Data+string(jsonData),
					out.PostbackActionData.Text,
					fmt.Sprintf(out.PostbackActionData.DisplayText, genre.Name),
					"",
					out.PostbackActionData.FillInText,
				),
			).WithImageOptions("#FFFFFF")
			ccs = append(ccs, cc)
		}
	}

	messages := linebot.NewTemplateMessage(
		out.TemplateMessageData.AltText,
		linebot.NewCarouselTemplate(ccs...).WithImageOptions("rectangle", "cover"),
	)

	if _, err := p.bot.ReplyMessage(out.ReplyToken, messages).Do(); err != nil {
		logrus.Errorf("Error LINEBOT replying message: %v", err)
	}

	return nil
}

func (p *LinePresenter) SetConditionMenuReplyMessage(out menudto.SetConditionMenuOutput) error {
	var firstCcs []*linebot.CarouselColumn
	// var remainCcs []*linebot.CarouselColumn

	// 最小の10個までを表示
	for _, condition := range out.Conditions {
		actions := []linebot.TemplateAction{}

		for _, option := range condition.ConditionOptions {
			newPostbackData := model.SetPostbackDataField(out.PostbackData, option.ParamKey, option.ParamValue)
			jsonData, err := json.Marshal(newPostbackData)
			if err != nil {
				logrus.Errorf("Error json.Marshal(newPostbackData): %v", err)
				return err
			}
			// ConditionOptionごとにアクションを作成
			action := linebot.NewPostbackAction(
				option.Name,
				out.PostbackActionData.Data+string(jsonData),
				out.PostbackActionData.Text,
				fmt.Sprintf(out.PostbackActionData.DisplayText, condition.Name, model.PBA_LABEL_CONDITION_SET_CONFIRM_MENU),
				"",
				out.PostbackActionData.FillInText,
			)
			actions = append(actions, action)
		}

		cc := linebot.NewCarouselColumn(
			condition.ImgURL,
			condition.Name,
			condition.Desc,
			actions...,
		).WithImageOptions("#FFFFFF")
		firstCcs = append(firstCcs, cc)
	}

	// 11個目以降を表示
	// for _, condition := range out.Conditions[10:] {
	// 	actions := []linebot.TemplateAction{}
	// 	for _, option := range condition.ConditionOptions {
	// 		// ConditionOptionごとにアクションを作成
	// 		action := linebot.NewPostbackAction(
	// 			option.Name,
	// 			option.ParamKey,
	// 			"",
	// 			option.Name,
	// 			"",
	// 			"")
	// 		actions = append(actions, action)
	// 	}

	// 	cc := linebot.NewCarouselColumn(
	// 		condition.ImgURL,
	// 		condition.Name,
	// 		condition.Desc,
	// 		actions...,
	// 	).WithImageOptions("#FFFFFF")
	// 	remainCcs = append(remainCcs, cc)
	// }

	var sms []linebot.SendingMessage

	firstTm := linebot.NewTemplateMessage("条件を選択してください", linebot.NewCarouselTemplate(firstCcs...).WithImageOptions("rectangle", "cover"))

	// remainTm := linebot.NewTemplateMessage("条件を選択してください", linebot.NewCarouselTemplate(remainCcs...).WithImageOptions("rectangle", "cover"))

	// sms = append(sms, firstTm, remainTm)
	sms = append(sms, firstTm)

	if _, err := p.bot.ReplyMessage(out.ReplyToken, sms...).Do(); err != nil {
		logrus.Errorf("Error LINEBOT replying message: %v", err)
		return err
	}
	return nil
}

func (p *LinePresenter) SetRestaurantListReplyMessage(out searchdto.Output) error {
	var bcs []*linebot.BubbleContainer
	for _, s := range out.Response.Results.Shops {
		b := linebot.BubbleContainer{
			Type:   linebot.FlexContainerTypeBubble,
			Hero:   setHero(s),
			Body:   setBody(s),
			Footer: setFooter(s, &out),
		}
		bcs = append(bcs, &b)
	}

	cc := &linebot.CarouselContainer{
		Type:     linebot.FlexContainerTypeCarousel,
		Contents: bcs,
	}

	messages := linebot.NewFlexMessage(out.TemplateMessageData.AltText, cc)
	if _, err := p.bot.ReplyMessage(out.ReplyToken, messages).Do(); err != nil {
		logrus.Errorf("Error LINEBOT replying message: %v", err)
		return err
	}
	return nil
}

// ヒーロー
func setHero(s *model.Shop) *linebot.ImageComponent {
	if s.Photo.Mobile.L == "" {
		return nil
	}
	return &linebot.ImageComponent{
		Type:        linebot.FlexComponentTypeImage,
		URL:         s.Photo.Mobile.L,
		Size:        linebot.FlexImageSizeTypeFull,
		AspectRatio: linebot.FlexImageAspectRatioType20to13,
		AspectMode:  linebot.FlexImageAspectModeTypeCover,
	}
}

func (p *LinePresenter) SetNothingHitReplyMessage(out searchdto.Output) error {
	if _, err := p.bot.ReplyMessage(out.ReplyToken, linebot.NewTextMessage(out.TextMessageData.Content)).Do(); err != nil {
		logrus.Errorf("Error LINEBOT replying message: %v", err)
		return err
	}
	return nil
}

// ボディ
func setBody(s *model.Shop) *linebot.BoxComponent {

	// ジャンル - タグコンポーネント
	gc := &linebot.BoxComponent{
		Type:            linebot.FlexComponentTypeBox,
		Layout:          linebot.FlexBoxLayoutTypeVertical,
		Position:        linebot.FlexComponentPositionTypeAbsolute,
		Width:           "85px",
		Height:          "20px",
		BackgroundColor: "#ff334b",
		CornerRadius:    "lg",
		OffsetTop:       "1.5%",
		OffsetStart:     "4%",
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   utils.ReplaceEmptyWithUnknown(s.Genre.Name),
				Margin: "3px",
				Size:   linebot.FlexTextSizeTypeXxs,
				Align:  linebot.FlexComponentAlignTypeCenter,
				Color:  "#ffffff",
			},
		},
	}

	// 喫煙 - タグコンポーネント
	sc := &linebot.BoxComponent{
		Type:            linebot.FlexComponentTypeBox,
		Layout:          linebot.FlexBoxLayoutTypeVertical,
		Position:        linebot.FlexComponentPositionTypeAbsolute,
		Width:           "85px",
		Height:          "20px",
		BackgroundColor: "#996633",
		CornerRadius:    "lg",
		OffsetTop:       "1.5%",
		OffsetStart:     "39%",
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   utils.ReplaceEmptyWithUnknown(s.NonSmoking),
				Margin: "3px",
				Size:   linebot.FlexTextSizeTypeXxs,
				Align:  linebot.FlexComponentAlignTypeCenter,
				Color:  "#ffffff",
			},
		},
	}

	// 予算 - タグコンポーネント
	bc := &linebot.BoxComponent{
		Type:            linebot.FlexComponentTypeBox,
		Layout:          linebot.FlexBoxLayoutTypeVertical,
		Position:        linebot.FlexComponentPositionTypeAbsolute,
		Width:           "100px",
		Height:          "20px",
		BackgroundColor: "#CCCC33",
		CornerRadius:    "lg",
		OffsetTop:       "1.5%",
		OffsetStart:     "74%",
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:   linebot.FlexComponentTypeText,
				Text:   utils.ReplaceEmptyWithUnknown(s.Budget.Name),
				Margin: "3px",
				Size:   linebot.FlexTextSizeTypeXxs,
				Align:  linebot.FlexComponentAlignTypeCenter,
				Color:  "#ffffff",
			},
		},
	}

	// セパレートコンポーネント
	dc := &linebot.SeparatorComponent{
		Type:   linebot.FlexComponentTypeSeparator,
		Margin: linebot.FlexComponentMarginTypeLg,
	}

	// タイトルコンポーネント
	tc := &linebot.TextComponent{
		Type:   linebot.FlexComponentTypeText,
		Text:   s.Name,
		Wrap:   true,
		Weight: linebot.FlexTextWeightTypeBold,
		Size:   linebot.FlexTextSizeTypeXl,
	}

	// アクセス - 詳細コンポーネント
	ac := setDetail("アクセス", utils.ReplaceEmptyWithUnknown(s.Access))

	// キャッチ - 詳細コンポーネント
	cc := setDetail("キャッチ", utils.ReplaceEmptyWithUnknown(s.Catch+"\n"+s.Genre.Catch))

	// 時間 - 詳細コンポーネント
	tic := setDetail("時間", utils.ReplaceEmptyWithUnknown(s.Open))

	// 定休日 - 詳細コンポーネント
	clc := setDetail("定休日", utils.ReplaceEmptyWithUnknown(s.Close))

	return &linebot.BoxComponent{
		Type:    linebot.FlexComponentTypeBox,
		Layout:  linebot.FlexBoxLayoutTypeVertical,
		Spacing: linebot.FlexComponentSpacingTypeMd,
		Contents: []linebot.FlexComponent{
			dc,  // セパレートコンポーネント
			tc,  // タイトルコンポーネント
			ac,  // アクセス - 詳細コンポーネント
			cc,  // キャッチ - 詳細コンポーネント
			tic, // 時間 - 詳細コンポーネント
			clc, // 定休日 - 詳細コンポーネント
			gc,  // ジャンル - タグコンポーネント
			sc,  // 喫煙 - タグコンポーネント
			bc,  // 予算 - タグコンポーネント
		},
	}
}

// フッター
func setFooter(s *model.Shop, o *searchdto.Output) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:    linebot.FlexComponentTypeBox,
		Layout:  linebot.FlexBoxLayoutTypeVertical,
		Spacing: linebot.FlexComponentSpacingTypeXs,
		Contents: []linebot.FlexComponent{
			setButton(o.URIActionDataForMore.Label, s.URLS.PC),
			setButton(o.URIActionDataForMap.Label, fmt.Sprintf(o.URIActionDataForMap.URI, url.QueryEscape(s.Name))),
		},
	}
}

// 詳細コンポーネントをセット
func setDetail(subTitle string, content string) *linebot.BoxComponent {
	return &linebot.BoxComponent{
		Type:    linebot.FlexComponentTypeBox,
		Layout:  linebot.FlexBoxLayoutTypeBaseline,
		Margin:  linebot.FlexComponentMarginTypeLg,
		Spacing: linebot.FlexComponentSpacingTypeSm,
		Contents: []linebot.FlexComponent{
			setSubtitle(subTitle),
			setContent(content),
		},
	}
}

// サブタイトルをセット
func setSubtitle(t string) *linebot.TextComponent {
	return &linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  t,
		Color: "#aaaaaa",
		Size:  linebot.FlexTextSizeTypeXs,
		Flex:  linebot.IntPtr(4),
	}
}

// 内容文をセット
func setContent(t string) *linebot.TextComponent {
	return &linebot.TextComponent{
		Type:  linebot.FlexComponentTypeText,
		Text:  t,
		Wrap:  true,
		Color: "#666666",
		Size:  linebot.FlexTextSizeTypeXs,
		Flex:  linebot.IntPtr(12),
	}
}

// ボタンをセット
func setButton(label string, uri string) *linebot.ButtonComponent {
	return &linebot.ButtonComponent{
		Type:   linebot.FlexComponentTypeButton,
		Style:  linebot.FlexButtonStyleTypeLink,
		Height: linebot.FlexButtonHeightTypeSm,
		Action: linebot.NewURIAction(label, uri),
	}
}
