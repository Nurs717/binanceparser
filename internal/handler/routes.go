package handler

import (
	"net/http"
)

func SetUpRoutes(mux *http.ServeMux) {

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", IndexHandler)
	mux.HandleFunc("/ws", WebSocketHandler)
}
