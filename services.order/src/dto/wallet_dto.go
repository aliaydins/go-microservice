package dto

type WalletDto struct {
	Wallet struct {
		ID     int `json:"id"`
		UserId int `json:"user_id"`
		USD    int `json:"usd"`
		BTC    int `json:"btc"`
	} `json:"wallet"`
}
