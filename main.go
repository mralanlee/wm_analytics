package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// websockets implementation
	http.HandleFunc("/payments", payments)

	// if this works out better, remove the web sockets implementation
	http.HandleFunc("/capture", capturePayments)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
