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

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// upgrades http connection to websocket connection
	ws, err := Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}
	var Ch = make(chan client.Sums)
	client.Conns[ws] = Ch
	// writes data from channel to client websocket connection
	go Writer(ws)
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	// upgrading http connection to websocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error occured when upgrading websocket connection %v\n", err)
		return ws, err
	}
	return ws, nil
}

func Writer(conn *websocket.Conn) {
	// initializing Sums struct
	var sums client.Sums
	for {
		// reading from channel to variable
		sums = <-client.Conns[conn]

		// marshalling data and sending to client
		err := conn.WriteJSON(sums)
		if err != nil {
			// deletes than kills connection if can't marshall
			delete(client.Conns, conn)
			log.Printf("error occured when marshalling websocket %v\n", err)
			conn.Close()
			break
		}
	}
}
