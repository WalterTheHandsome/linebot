package ccboybotai

import (
	"log"
	"os"

	"github.com/WalterTheHandsome/linebot/gemini"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

var (
	aiAPIKey           string
	channelSecret      string
	channelAccessToken string
	bot                *messaging_api.MessagingApiAPI
	blob               *messaging_api.MessagingApiBlobAPI
	userID             string
)

func Init() {
	aiAPIKey = os.Getenv(ENV_AI_API_KEY)
	channelSecret = os.Getenv(ENV_AI_BOT_CHANNEL_SECRET)
	channelAccessToken = os.Getenv(ENV_AI_BOT_CHANNEL_ACCESS_TOKEN)
	userID = os.Getenv(ENV_AI_BOT_USER_ID)

	gemini.Init(aiAPIKey)
	var err error
	bot, err = messaging_api.NewMessagingApiAPI(channelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	blob, err = messaging_api.NewMessagingApiBlobAPI(channelAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("ccboyai Init() done")
}
