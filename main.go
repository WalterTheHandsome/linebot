package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/WalterTheHandsome/linebot/ccboybot"
	"github.com/WalterTheHandsome/linebot/ccboybotai"
)

func main() {
	isAI := os.Getenv("IS_AI")
	if isAI == "true" {
		log.Println("Init ccboybot ai")
		ccboybotai.Init()
		http.HandleFunc(ccboybotai.ROUTE_PATH, ccboybotai.MainHandler)
		http.HandleFunc(ccboybotai.ROUTE_AUTH, ccboybotai.AuthHandler)
	} else {
		ccboybot.Init()
		log.Println("Init ccboybot ")
		http.HandleFunc(ccboybot.ROUTE_PATH, ccboybot.MainHandler)
	}

	port := os.Getenv("SERVER_PORT") // for dev
	if port == "" {
		port = os.Getenv("PORT") // Heroku will assign a port randomly
	}
	fmt.Println("line bot server starts on ->", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
