package main

import (
	"cryptotest/internal/app"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/balance/{address}", app.Balance)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
