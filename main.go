package main

import (
	"log"
	"net/http"
	"os"

	"github.com/WalterTheHandsome/linebot/ccboybot"
)

const (
	DEFAULT_SERVER_PORT      = "80"
	ccboyCredentialPath      = "./ignored/ccboy-credential.yml"
	ccboyLoginCredentialPath = "./ignored/ccboy-login-credential.yml"
)

func main() {
	ccboybot.Init(ccboyCredentialPath)
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc(ccboybot.ROUTE_PATH, ccboybot.MainHandler)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = DEFAULT_SERVER_PORT
	}
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
