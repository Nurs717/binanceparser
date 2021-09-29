package client

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

type Data struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

type Sums struct {
	SumAsks float64
	SumBids float64
}

func GetData(ch chan<- Sums) {
	ws, _, err := websocket.DefaultDialer.Dial("wss://stream.binance.com:9443/ws/btcusdt@depth20@1000ms", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	reader(ws, ch)
}

func reader(conn *websocket.Conn, ch chan<- Sums) {
	var data Data
	for {
		err := conn.ReadJSON(&data)
		if err != nil {
			log.Println(err)
		}

		sum := Sum(&data)
		fmt.Println(sum)

		ch <- sum
	}
}

func Sum(item *Data) Sums {
	var sum Sums
	for i := 0; i < 15; i++ {
		ask, err := strconv.ParseFloat(item.Asks[i][1], 64)
		if err != nil {
			log.Println(err)
		}
		bid, err := strconv.ParseFloat(item.Bids[i][1], 64)
		if err != nil {
			log.Println(err)
		}
		sum.SumAsks += ask
		sum.SumBids += bid
	}
	return sum
}
