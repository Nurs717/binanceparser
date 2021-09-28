package model

import (
	"log"

	"github.com/gorilla/websocket"
)

type Item struct {
	ID   string   `json:"lastUpdateId"`
	Bids []string `json:"bids"`
	Asks []string `json:"asks"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func GetData() *Item {
	var item *Item
	ws, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/btcusdt@depth20@1000ms", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	reader(ws)
	// err = json.Unmarshal(data, &item)
	// if err != nil {
	// 	log.Println(err)
	// }
	return item
}

func reader(conn *websocket.Conn) {
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
		}

		// return p

		log.Println(string(p))
		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Println(err)
		// 	return
		// }
	}
}
