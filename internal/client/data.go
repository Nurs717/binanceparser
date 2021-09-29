package client

import (
	"log"
	"strconv"

	"github.com/gorilla/websocket"
)

//struct for data comming from binance websocket
type Data struct {
	Bids [][]string `json:"bids"`
	Asks [][]string `json:"asks"`
}

//struct to calculate sum of asks and bids
type Sums struct {
	SumAsks float64
	SumBids float64
}

const binanceUrl = "wss://stream.binance.com:9443/ws/btcusdt@depth20@1000ms"

func GetData(ch chan<- Sums) {
	// sets connection with Binance websocket
	ws, _, err := websocket.DefaultDialer.Dial(binanceUrl, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	reader(ws, ch)
}

func reader(conn *websocket.Conn, ch chan<- Sums) {
	//initializing struct
	var data Data
	for {
		//unmarshaling websocket connection
		err := conn.ReadJSON(&data)
		if err != nil {
			log.Printf("error occured when unmarshaling websocket %v\n", err)
		}

		// calculates total sum of first 15 postion of asks and bids
		// and writes results to Sums struct
		// and return Sums struct
		sum := Sum(&data)
		log.Printf("Ask Orders: %v Bids Orders: %v\n", sum.SumAsks, sum.SumBids)

		// writes struct Sums to channel
		ch <- sum
	}
}

func Sum(item *Data) Sums {
	//initializing struct
	var sum Sums
	for i := 0; i < 15; i++ {
		ask, err := strconv.ParseFloat(item.Asks[i][1], 64)
		if err != nil {
			log.Printf("error occured when converting string to float64 %v\n", err)
		}
		bid, err := strconv.ParseFloat(item.Bids[i][1], 64)
		if err != nil {
			log.Printf("error occured when converting string to float64 %v\n", err)
		}
		sum.SumAsks += ask
		sum.SumBids += bid
	}
	return sum
}