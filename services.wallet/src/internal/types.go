package wallet

import "github.com/aliaydins/microservice/service.wallet/src/entity"

type WalletDTO struct {
	UserId int `json:"user_id"`
	USD    int `json:"usd"`
	BTC    int `json:"btc"`
}

func mapper(wallet *entity.Wallet) *WalletDTO {
	dto := &WalletDTO{
		UserId: wallet.UserId,
		USD:    wallet.USD,
		BTC:    wallet.BTC,
	}

	return dto
}
