package ccboybotai

import (
	"log"

	"github.com/WalterTheHandsome/linebot/gemini"
	"github.com/WalterTheHandsome/linebot/lib"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var (
	aiAPIKey      string
	channelSecret string
	bot           *messaging_api.MessagingApiAPI
	blob          *messaging_api.MessagingApiBlobAPI
	userID        string
)

func Init(credentialPath string) {
	credentials, err := lib.ReadCredential(credentialPath)
	if err != nil {
		log.Fatal(err)
	}
	channelSecret = credentials.LineChannelSecret
	userID = credentials.UserID
	aiAPIKey = credentials.AIAPIKey
	gemini.Init(aiAPIKey)

	bot, err = messaging_api.NewMessagingApiAPI(credentials.LineChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	blob, err = messaging_api.NewMessagingApiBlobAPI(credentials.LineChannelAccessToken)
	if err != nil {
		log.Fatal(err)
	}
}
