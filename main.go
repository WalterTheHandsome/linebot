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
	ccboybotai.Init()
	ccboybot.Init()

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc(ccboybot.ROUTE_PATH, ccboybot.MainHandler)

	http.HandleFunc(ccboybotai.ROUTE_PATH, ccboybotai.MainHandler)
	http.HandleFunc(ccboybotai.ROUTE_AUTH, ccboybotai.AuthHandler)

	port := os.Getenv("SERVER_PORT") // for dev
	if port == "" {
		port = os.Getenv("PORT") // heroku will assign a port randomly
	}
	fmt.Println("line bot server starts on ->", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
