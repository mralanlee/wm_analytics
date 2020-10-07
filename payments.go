package main

import (
	"encoding/json"
	"net/http"

	"github.com/mralanlee/wm_analytics/consumers"
)

func capturePayments(w http.ResponseWriter, r *http.Request) {
	var p consumers.Payments

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&p); err != nil {
		http.Error(w, "Error reading request", http.StatusInternalServerError)
		return
	}

	for _, v := range p.Payments {
		go func(detail consumers.PaymentDetail) {
			detail.Transform()
			result := consumers.PostgresClient.Create(detail)

			if result.Error != nil {
				http.Error(w, "Error processing transaction", http.StatusInternalServerError)
				return
			}
		}(v)
	}

	w.WriteHeader(200)
}
