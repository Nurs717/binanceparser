package main

import (
	"binanceparser/internal/handler"
	"binanceparser/internal/model"
	"binanceparser/internal/server"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler.SetUpRoutes(mux)

	fmt.Println("Starting at localhost:8080")
	srv := server.NewServer("8081", mux)
	model.GetData()
	if err := srv.Run(); err != nil {
		log.Fatalf("error occured while running API server: %v", err)
	}
}
