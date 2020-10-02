package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mralanlee/wm_analytics/consumers"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func payments(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("Upgrade:", err)
		return
	}

	for {
		var raw json.RawMessage
		err := c.ReadJSON(&raw)
		if err != nil {
			log.Println("read:", err)
			break
		}

		var payDetails consumers.PaymentDetail

		jsonErr := json.Unmarshal(raw, &payDetails)

		if jsonErr != nil {
			log.Println("json parse:", jsonErr)
		}

		payDetails.Transform()

		consumers.PostgresClient.Create(&payDetails)
	}
}
