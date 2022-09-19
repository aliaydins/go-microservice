package event

type OrderCompleted struct {
	UserId int    `json:"user_id"`
	Type   string `json:"type"`
	USD    int    `json:"usd"`
	BTC    int    `json:"btc"`
	Amount int    `json:"amount"`
}
