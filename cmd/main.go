package main

import (
	"cryptotest/internal/app"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/balance/{address}", balance)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func balance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("новый запрос")
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[len("/balance/"):])

	// Получаем адрес кошелька из URL
	address := r.URL.Path[len("/balance/"):]

	// Получаем баланс кошелька
	balance, err := app.GetBalance(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Возвращаем баланс в ответе
	fmt.Fprintf(w, "%s", balance)
}
