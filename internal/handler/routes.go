package handler

import (
	"net/http"
)

func SetUpRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/ws", WebSocketHandler)
}
