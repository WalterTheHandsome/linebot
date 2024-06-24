package ccboybot

import (
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

type ToUX struct {
	Range string
	URL   string
}

type ToIT struct {
	Range   string
	Content string
	URL     string
}

const (
	ToUXTemplate = "%s的切版更新到demo機了,再麻煩你有空的時候幫忙驗收,感謝\n%s"
	ToITTemplate = "下方為%s的切版檔,%s已更新上測試機,再麻煩了,謝謝~\n%s"

	LionMainMenuCarouselJsonString = `{
		"type": "carousel",
		"contents": [
			{
				"type": "bubble",
				"body": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "–切版檔UX確認文案–"
						}
					]
				},
				"footer": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "button",
							"action": {
								"type": "message",
								"label": "開始",
								"text": "/toUX"
							}
						}
					]
				}
			},
			{
				"type": "bubble",
				"body": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "text",
							"text": "–切版檔交付IT文案–"
						}
					]
				},
				"footer": {
					"type": "box",
					"layout": "vertical",
					"contents": [
						{
							"type": "button",
							"action": {
								"type": "message",
								"label": "開始",
								"text": "/toIT"
							}
						}
					],
					"action": {
						"type": "message",
						"label": "開始",
						"text": "/toIT"
					}
				}
			}
		]
	}`
)

func ReplyLionCarousel(replyToken string) {
	contents, err := messaging_api.UnmarshalFlexContainer([]byte(LionMainMenuCarouselJsonString))
	if err != nil {
		log.Println("parse carousel err", err)
	}
	if _, err := bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.FlexMessage{
					AltText:  "Flex message alt text",
					Contents: contents,
				},
			},
		},
	); err != nil {
		log.Println("reply message err", err)
	}
}

var (
	toUX = new(ToUX)
	toIT = new(ToIT)
)

func reset() {
	botState = STATE_NONE
	toUX = new(ToUX)
	toIT = new(ToIT)
}

func HandleMenuLion(message webhook.TextMessageContent, event webhook.MessageEvent) {
	msg := message.Text

	switch botState {
	case STATE_LINE_PENDING_FOR_CHOOSE:
		if IsCommand(msg) {
			switch msg {
			case MENU_LION_TO_UX:
				botState = STATE_LION_PENDING_FOR_TO_UX_RANGE
			case MENU_LION_TO_IT:
				botState = STATE_LION_PENDING_FOR_TO_IT_RANGE
			}
			ReplyTextMessage("請輸入範圍", event.ReplyToken)
			return
		}
	case STATE_LION_PENDING_FOR_TO_UX_RANGE:
		if !IsCommand(msg) {
			botState = STATE_LION_PENDING_FOR_TO_UX_URL
			toUX.Range = msg
			ReplyTextMessage("請輸入連結", event.ReplyToken)
			return
		}
	case STATE_LION_PENDING_FOR_TO_UX_URL:
		if !IsCommand(msg) {
			toUX.URL = msg
			reply := fmt.Sprintf(ToUXTemplate, toUX.Range, toUX.URL)
			ReplyTextMessage(reply, event.ReplyToken)
			reset()
			return
		}
	case STATE_LION_PENDING_FOR_TO_IT_RANGE:
		if !IsCommand(msg) {
			toIT.Range = msg
			botState = STATE_LION_PENDING_FOR_TO_IT_CONTENT
			ReplyTextMessage("請輸入更新內容", event.ReplyToken)
			return
		}
	case STATE_LION_PENDING_FOR_TO_IT_CONTENT:
		if !IsCommand(msg) {
			toIT.Content = msg
			botState = STATE_LION_PENDING_FOR_TO_IT_URL
			ReplyTextMessage("請輸入連結", event.ReplyToken)
			return
		}
	case STATE_LION_PENDING_FOR_TO_IT_URL:
		if !IsCommand(msg) {
			toIT.URL = msg
			reply := fmt.Sprintf(ToITTemplate, toIT.Range, toIT.Content, toIT.URL)
			ReplyTextMessage(reply, event.ReplyToken)
			reset()
			return
		}
	default:
		reset()
	}
}
