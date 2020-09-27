package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)
	http.HandleFunc("/payments", payments)
	err := http.ListenAndServe(":3000", nil)

	if err != nil {
		log.Fatal(err)
	}
}
