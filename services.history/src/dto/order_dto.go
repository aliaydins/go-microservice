package dto

type OrderDTO struct {
	UserId int `json:"user_id"`
	Type   string
	USD    int
	BTC    int
	Amount int
}
