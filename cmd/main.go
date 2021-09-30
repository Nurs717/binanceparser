package main

import (
	"binanceparser/internal/client"
	"binanceparser/internal/handler"
	"binanceparser/internal/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	// set handlers and static file
	handler.SetUpRoutes(mux)

	//set server configs
	srv := server.NewServer("8080", mux)

	// gets asks and bids from binance api
	// and writes to all connection channels
	go client.GetData()

	// runs server
	fmt.Println("Starting at localhost:8080")
	if err := srv.Run(); err != nil {
		log.Fatalf("error occured while running API server: %v", err)
	}
}
