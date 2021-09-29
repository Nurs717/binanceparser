package handler

import (
	"binanceparser/internal/client"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var Ch = make(chan client.Sums)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	go Writer(ws, Ch)
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}
	return ws, nil
}

func Writer(conn *websocket.Conn, ch <-chan client.Sums) {
	var sums client.Sums
	for {
		sums = <-ch

		err := conn.WriteJSON(sums)
		if err != nil {
			log.Println(err)
		}
	}
}
