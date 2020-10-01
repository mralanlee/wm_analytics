package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


type PaymentDetail struct {
	UrlPath string `json:"urlPath"`
	PaymentPointer string `json:"paymentPointer"`
	RequestID string `json:"requestId"`
	Amount string `json:"amount"`
	AssetCode string `json:"assetCode"`
	AssetScale int `json:"assetScale"`
	Receipt string `json:"receipt"`
}


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

		var payDetails PaymentDetail

		jsonErr := json.Unmarshal(raw, &payDetails)

		if jsonErr != nil {
			log.Println("json parse:", jsonErr)
		}

		log.Println(payDetails)
	}
}
