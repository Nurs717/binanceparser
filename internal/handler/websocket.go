package handler

import (
	"fmt"
	"net/http"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "WebSocket Endpoint")
	// upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// ws, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	log.Println(err)
	// }
	// log.Println("WebSocket connected")
	// reader(ws)
}
