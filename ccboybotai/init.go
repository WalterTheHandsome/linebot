package ccboybotai

import (
	"log"

	"github.com/WalterTheHandsome/linebot/lib"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var (
	channelSecret string
	bot           *messaging_api.MessagingApiAPI
	userID        string
)

func Init(credentialPath string) {
	credentials, err := lib.ReadCredential(credentialPath)
	if err != nil {
		log.Fatal(err)
	}
	channelSecret = credentials.LineChannelSecret
	userID = credentials.UserID
	bot, err = messaging_api.NewMessagingApiAPI(
		credentials.LineChannelAccessToken,
	)
	if err != nil {
		log.Fatal(err)
	}
}
