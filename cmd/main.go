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
	handler.SetUpRoutes(mux)
	srv := server.NewServer("8080", mux)

	go client.GetData(handler.Ch)

	fmt.Println("Starting at localhost:8080")
	if err := srv.Run(); err != nil {
		log.Fatalf("error occured while running API server: %v", err)
	}
}
