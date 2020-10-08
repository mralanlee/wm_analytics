package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mralanlee/wm_analytics/consumers"
)

func capturePayments(w http.ResponseWriter, r *http.Request) {
	log.Println("Incoming request")
	if r.Method != http.MethodPost {
		log.Fatal("Invalid http method")
		http.Error(w, "Request method not supported", http.StatusMethodNotAllowed)
		return
	}

	var p consumers.Payments

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&p); err != nil {
		log.Fatal(err)
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	for _, v := range p.Details {
		go func(detail *consumers.PaymentDetail) {
			detail.Transform()
			result := consumers.PostgresClient.Create(detail)

			if result.Error != nil {
				log.Fatal(result.Error)
				http.Error(w, "Error processing transaction", http.StatusInternalServerError)
				return
			}
		}(&v)
	}

	w.WriteHeader(200)
}
