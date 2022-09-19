package history

import "github.com/aliaydins/microservice/service.history/src/entity"

type HistoryDTO struct {
	UserId int    `json:"user_id"`
	Type   string `json:"type"`
	USD    int    `json:"usd"`
	BTC    int    `json:"btc"`
	Amount int    `json:"amount"`
}

func mapper(history entity.History) HistoryDTO {
	dto := HistoryDTO{
		UserId: history.UserId,
		Type:   history.Type,
		USD:    history.USD,
		BTC:    history.BTC,
		Amount: history.Amount,
	}

	return dto
}
