package ccboybot

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func PushTextMessage(msg string, user string) {
	pushMsg := new(messaging_api.PushMessageRequest)
	pushMsg.To = user
	pushMsg.Messages = []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: msg,
		},
	}
	bot.PushMessage(pushMsg, "")
}

func ReplyTextMessage(reply string, replyToken string) {
	if _, err := bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				messaging_api.TextMessage{
					Text: reply,
				},
			},
		},
	); err != nil {
		log.Println("reply message err->", err)
	} else {
		log.Println("Sent text reply.")
	}
}

func IsCommand(from string) bool {
	return strings.HasPrefix(from, "/")
}

func TextMessageRouter(message webhook.TextMessageContent, event webhook.MessageEvent) {
	msg := message.Text
	if !IsCommand(msg) && botState == STATE_NONE {
		return
	}
	if botState == STATE_NONE {
		switch msg {
		case MENU_LION:
			botState = STATE_LION_PENDING_FOR_CHOOSE
			ReplyLionCarousel(event.ReplyToken)
			return
		}
	}

	if botState > STATE_NONE && botState <= STATE_LION_PENDING_FOR_TO_ONLINE_2_AND_3 {
		HandleMenuLion(message, event)
		return
	}

	ReplyTextMessage("none", event.ReplyToken)
}

func MainHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s called...\n", ROUTE_PATH)

	cb, err := webhook.ParseRequest(channelSecret, req)
	if err != nil {
		log.Printf("Cannot parse request: %+v\n", err)
		if errors.Is(err, webhook.ErrInvalidSignature) {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	log.Println("Handling events...")
	for _, event := range cb.Events {
		log.Printf("%s called%+v...\n", ROUTE_PATH, event)

		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				TextMessageRouter(message, e)
			case webhook.StickerMessageContent:
				replyMessage := fmt.Sprintf(
					"sticker id is %s, stickerResourceType is %s, stickerPackageId is %s", message.StickerId, message.StickerResourceType, message.PackageId)
				fmt.Println("reply message is", replyMessage)
				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: replyMessage,
							},
							messaging_api.StickerMessage{
								StickerId: message.StickerId,
								PackageId: message.PackageId,
							},
						},
					}); err != nil {
					log.Print("reply msg err ->", err)
				} else {
					log.Println("Sent sticker reply.")
				}
			default:
				log.Printf("Unsupported message content: %T\n", e.Message)
			}
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
}
