package app

type BalanceResponse struct {
	Address string `json:"address"`
	Balance string `json:"balance"`
}

func GetBalance(Address string) (string, error) {
	return "800", nil
}
