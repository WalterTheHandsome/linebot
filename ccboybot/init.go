package ccboybot

import (
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var (
	channelSecret      string
	userID             string
	channelAccessToken string

	bot      *messaging_api.MessagingApiAPI
	botState int
)

func Init() {
	channelSecret = os.Getenv(ENV_BOT_CHANNEL_SECRET)
	channelAccessToken = os.Getenv(ENV_BOT_CHANNEL_ACCESS_TOKEN)
	userID = os.Getenv(ENV_BOT_USER_ID)
	var err error
	bot, err = messaging_api.NewMessagingApiAPI(channelAccessToken)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ccboy Init() done")
}
