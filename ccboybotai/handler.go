package ccboybotai

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/WalterTheHandsome/linebot/gemini"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"

	"github.com/google/generative-ai-go/genai"
)

var (
	userSessions = make(map[string]*genai.ChatSession)
)

func ShowLoadingAnimation(chatID string, duration int32) {
	animationRequest := new(messaging_api.ShowLoadingAnimationRequest)
	animationRequest.ChatId = chatID
	animationRequest.LoadingSeconds = duration
	bot.ShowLoadingAnimation(animationRequest)
}

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

func AuthHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("req is", req.Body)
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
				HanleTextMessageContent(message, e)
			case webhook.StickerMessageContent:
				HandleStickerMessageContent(message, e.ReplyToken)
			case webhook.ImageMessageContent:
				log.Println("Got img msg ID:", message.Id)

				//Get image binary from LINE server based on message ID.
				content, err := blob.GetMessageContent(message.Id)
				if err != nil {
					log.Println("Got GetMessageContent err:", err)
				}
				defer content.Body.Close()
				data, err := io.ReadAll(content.Body)
				if err != nil {
					log.Fatal(err)
				}
				ret, err := gemini.GeminiImage(data)
				if err != nil {
					ret = "無法辨識圖片內容，請重新輸入:" + err.Error()
				}
				if err := replyText(e.ReplyToken, ret); err != nil {
					log.Print(err)
				}
			case webhook.VideoMessageContent:
			case webhook.AudioMessageContent:
			case webhook.LocationMessageContent:
			default:
				log.Printf("Unsupported message content: %T\n", e.Message)
			}
		case webhook.FollowEvent:
			log.Printf("message: Got followed event")
		case webhook.PostbackEvent:
			data := e.Postback.Data
			log.Printf("Unknown message: Got postback: " + data)
		case webhook.BeaconEvent:
			log.Printf("Got beacon: " + e.Beacon.Hwid)
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
}

func replyText(replyToken, text string) error {
	if _, err := bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
			Messages: []messaging_api.MessageInterface{
				&messaging_api.TextMessage{
					Text: text,
				},
			},
		},
	); err != nil {
		log.Println(err)
		return err
	}
	log.Println("Sent text reply.")
	return nil
}

func HanleTextMessageContent(message webhook.TextMessageContent, event webhook.MessageEvent) {
	if !strings.HasPrefix(message.Text, "/") {
		log.Println("Wont response -> ", message.Text)
		return
	}

	req := message.Text
	// 檢查是否已經有這個用戶的 ChatSession or req == "reset"

	// 取得用戶 ID
	var uID string
	switch source := event.Source.(type) {
	case *webhook.UserSource:
		uID = source.UserId
		ShowLoadingAnimation(uID, 15)
	case *webhook.GroupSource:
		uID = source.UserId
	case *webhook.RoomSource:
		uID = source.UserId
	}

	// 檢查是否已經有這個用戶的 ChatSession
	cs, ok := userSessions[uID]
	if !ok {
		// 如果沒有，則創建一個新的 ChatSession
		cs = gemini.StartNewChatSession()
		userSessions[uID] = cs
	}
	if req == "reset" {
		// 如果需要重置記憶，創建一個新的 ChatSession
		cs = gemini.StartNewChatSession()
		userSessions[uID] = cs
		if err := replyText(event.ReplyToken, "很高興初次見到你，請問有什麼想了解的嗎？"); err != nil {
			log.Print(err)
		}
		return
	}
	// 使用這個 ChatSession 來處理訊息 & Reply with Gemini result
	res := gemini.Send(cs, req)
	ret := gemini.PrintResponse(res)
	if err := replyText(event.ReplyToken, ret); err != nil {
		log.Print(err)
	}

}

func HandleStickerMessageContent(message webhook.StickerMessageContent, replyToken string) {
	replyMessage := fmt.Sprintf(
		"sticker id is %s, stickerResourceType is %s, stickerPackageId is %s",
		message.StickerId,
		message.StickerResourceType,
		message.PackageId)
	fmt.Println("reply message is", replyMessage)
	if _, err := bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: replyToken,
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
		log.Print(err)
	} else {
		log.Println("Sent sticker reply.")
	}
}
