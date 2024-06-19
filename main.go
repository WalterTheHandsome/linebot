package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WalterTheHandsome/linebot/ccboybot"
	"github.com/WalterTheHandsome/linebot/ccboybotai"
)

const (
	DEFAULT_SERVER_PORT = "80"
	// ccboyCredentialPath   = "./ignored/ccboy-credential.yml"
	// ccboyAICredentialPath = "./ignored/ccboy-ai-credential.yml"
)

func main() {
	ccboybotai.Init()
	ccboybot.Init()

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc(ccboybot.ROUTE_PATH, ccboybot.MainHandler)

	http.HandleFunc(ccboybotai.ROUTE_PATH, ccboybotai.MainHandler)
	http.HandleFunc(ccboybotai.ROUTE_AUTH, ccboybotai.AuthHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = DEFAULT_SERVER_PORT
	}
	fmt.Println("line bot server starts on ->", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
