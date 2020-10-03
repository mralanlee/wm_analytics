package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	"github.com/mralanlee/wm_analytics/consumers"
)

var addr = flag.String("addr", "localhost:3000", "server address")

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/payments"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		panic(err)
	}
	defer c.Close()

	done := make(chan struct{})

	payment := &consumers.PaymentDetail{
		PaymentPointer: "$test.wmanalytics.com/testuser",
		RequestID: "6eb35aed-ac9d-464b-93af-ae8752649642",
		Amount: "135",
		AssetCode: "XRP",
		AssetScale: 9,
		Receipt: "",
	}

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-done:
			return
		case t := <-ticker.C:
			log.Println("ticker at ", t)
			err := c.WriteJSON(&payment)
			if err != nil {
				log.Fatal(err)
				return
			}
		case <- interrupt:
			log.Println(interrupt)
			return
		}
	}

}
