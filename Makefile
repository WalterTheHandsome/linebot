NGROK_TOKEN=$(shell cat ./ignored/ngrok_token.txt)
DEV_SERVER_PORT=80

NowDirName=$$(echo ${PWD} | awk -F '/' '{print $$NF}')
GoBuildEnv=GO111MODULE=on

env:
	echo "GO1.20 is required"
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

run: build
	SERVER_PORT=$(DEV_SERVER_PORT) ./dist/$(NowDirName)

run-ngrok:
	echo "ngrok required -> https://ngrok.com/download"
	ngrok config add-authtoken $(NGROK_TOKEN)
	ngrok http http://localhost:$(DEV_SERVER_PORT)/


