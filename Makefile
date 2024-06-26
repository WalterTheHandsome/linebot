NGROK_TOKEN=$(shell cat ./ignored/ngrok_token.txt)
AI_DEV_SERVER_PORT=8080
DEV_SERVER_PORT=8090

NowDirName=$$(echo ${PWD} | awk -F '/' '{print $$NF}')
GoBuildEnv=GO111MODULE=on

env:
	echo "GO1.21 is required"
	GO111MODULE=on go get -u	"github.com/line/line-bot-sdk-go/v8/linebot"

mod-tidy:
	@printf "[_] Run go mod tidy\r"
	@GO111MODULE=on GOSUMDB=off go mod tidy
	@printf "[v] Run go mod tidy\n"

clean:
	@printf "[_] Clear build files\r"
	@rm -f ./$(NowDirName)
	@printf "[v] Clear build files\n"

build: ./main.go clean mod-tidy
	@go version
	@printf "[_] Building binary\r"
	@$(GoBuildEnv) go build -o ./dist/$(NowDirName) $<
	@printf "[v] Building binary\n"

BOT_WITH_AI_API_KEY=$(shell cat ./ignored/ccboy-ai-api-key.txt)
BOT_WITH_AI_CHANNEL_SECRET=$(shell cat ./ignored/ccboy-ai-channel-secret.txt)
BOT_WITH_AI_ACCESS_TOKEN=$(shell cat ./ignored/ccboy-ai-access-token.txt)
BOT_WITH_AI_CHANNEL_USER_ID=$(shell cat ./ignored/ccboy-ai-user-id.txt)
BOT_WITH_AI_FIREBASE_CREDENTIAL=$(shell cat ./ignored/line-bot-play-6f3e99a3a12d.json)
BOT_WITH_AI_FIREBASE_URL=$(shell cat ./ignored/firebase_url.txt)

BOT_WITHOUT_AI_CHANNEL_SECRET=$(shell cat ./ignored/ccboy-channel-secret.txt)
BOT_WITHOUT_AI_CHANNEL_ACCESS_TOKEN=$(shell cat ./ignored/ccboy-access-token.txt)
BOT_WITHOUT_AI_CHANNEL_USER_ID=$(shell cat ./ignored/ccboy-user-id.txt)

run-ai-bot: build
	BOT_AI_API_KEY=$(BOT_WITH_AI_API_KEY) \
	BOT_AI_CHANNEL_SECRET=$(BOT_WITH_AI_CHANNEL_SECRET) \
	BOT_AI_CHANNEL_ACCESS_TOKEN=$(BOT_WITH_AI_ACCESS_TOKEN) \
	BOT_AI_CHANNEL_USER_ID=$(BOT_WITH_AI_CHANNEL_USER_ID) \
	IS_AI=true \
	GOOGLE_APPLICATION_CREDENTIALS='$(BOT_WITH_AI_FIREBASE_CREDENTIAL)' \
	FIREBASE_URL=$(BOT_WITH_AI_FIREBASE_URL) \
	SERVER_PORT=$(AI_DEV_SERVER_PORT) ./dist/$(NowDirName)

run-bot: build
	BOT_CHANNEL_SECRET=$(BOT_WITHOUT_AI_CHANNEL_SECRET) \
	BOT_CHANNEL_ACCESS_TOKEN=$(BOT_WITHOUT_AI_CHANNEL_ACCESS_TOKEN) \
	BOT_CHANNEL_USER_ID=$(BOT_WITHOUT_AI_CHANNEL_USER_ID) \
	GOOGLE_APPLICATION_CREDENTIALS='$(BOT_WITH_AI_FIREBASE_CREDENTIAL)' \
	FIREBASE_URL=$(BOT_WITH_AI_FIREBASE_URL) \
	IS_AI=false \
	SERVER_PORT=$(DEV_SERVER_PORT) ./dist/$(NowDirName)

run-ngrok:
	echo "ngrok required -> https://ngrok.com/download"
	ngrok start --all --config="./ignored/ngrok.yml"



