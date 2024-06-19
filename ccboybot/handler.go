package ccboybot

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
)

func PushTextMessage(msg string) {
	pushMsg := new(messaging_api.PushMessageRequest)
	pushMsg.To = userID
	pushMsg.Messages = []messaging_api.MessageInterface{
		messaging_api.TextMessage{
			Text: msg,
		},
	}
	bot.PushMessage(pushMsg, "")
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
				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: "[c]" + message.Text,
							},
						},
					},
				); err != nil {
					log.Println("reply message err->", err)
				} else {
					log.Println("Sent text reply.")
				}
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
