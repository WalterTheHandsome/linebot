NGROK_TOKEN=$(shell cat ./ignored/ngrok_token.txt)
DEV_SERVER_PORT=80

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

BOT_WITH_AI_ACCESS_TOKEN=$(shell cat ./ignored/ccboy-ai-access-token.txt)
BOT_WITH_AI_API_KEY=$(shell cat ./ignored/ccboy-ai-api-key.txt)
BOT_WITH_AI_CHANNEL_SECRET=$(shell cat ./ignored/ccboy-ai-channel-secret.txt)
BOT_WITH_AI_CHANNEL_USER_ID=$(shell cat ./ignored/ccboy-ai-user-id.txt)

BOT_CHANNEL_ACCESS_TOKEN=$(shell cat ./ignored/ccboy-access-token.txt)
BOT_CHANNEL_SECRET=$(shell cat ./ignored/ccboy-channel-secret.txt)
BOT_CHANNEL_USER_ID=$(shell cat ./ignored/ccboy-user-id.txt)

run: build
	BOT_AI_CHANNEL_ACCESS_TOKEN=$(BOT_WITH_AI_ACCESS_TOKEN) \
	BOT_AI_API_KEY=$(BOT_WITH_AI_API_KEY) \
	BOT_AI_CHANNEL_SECRET=$(BOT_WITH_AI_CHANNEL_SECRET) \
	BOT_AI_CHANNEL_USER_ID=$(BOT_WITH_AI_CHANNEL_USER_ID) \
	BOT_CHANNEL_SECRET=$(BOT_CHANNEL_SECRET) \
	BOT_CHANNEL_ACCESS_TOKEN=$(BOT_CHANNEL_ACCESS_TOKEN) \
	BOT_CHANNEL_USER_ID=$(BOT_CHANNEL_USER_ID) \
	./dist/$(NowDirName)

run-ngrok:
	echo "ngrok required -> https://ngrok.com/download"
	ngrok config add-authtoken $(NGROK_TOKEN)
	ngrok http http://localhost:$(DEV_SERVER_PORT)/


