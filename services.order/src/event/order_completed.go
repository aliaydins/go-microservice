package event

type OrderCompleted struct {
	BuyerUserId  int `json:"buyer_user_id"` // maker
	SellerUserId int `json:"seller_user_id"`
	USD          int `json:"usd"`
	BTC          int `json:"btc"`
}
