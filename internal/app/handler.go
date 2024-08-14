package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Balance(w http.ResponseWriter, r *http.Request) {
	fmt.Println("новый запрос")
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Path[len("/balance/"):])

	// Получаем адрес кошелька из URL
	address := r.URL.Path[len("/balance/"):]

	// Получаем баланс кошелька
	balance, err := GetBalance(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := BalanceResponse{
		Address: address,
		Balance: balance,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Возвращаем баланс в ответе
	fmt.Fprintf(w, "%s", jsonResponse)
}
