package wallet

import "github.com/aliaydins/microservice/service.wallet/src/entity"

type WalletDTO struct {
	ID       int             `json:"id""`
	UserId   int             `json:"user_id"`
	Currency entity.Currency `json:"currency"`
	Quantity int             `json:"quantity"`
}

func mapper(wallet *entity.Wallet) *WalletDTO {
	dto := &WalletDTO{
		ID:       wallet.ID,
		UserId:   wallet.UserId,
		Currency: wallet.Currency,
		Quantity: wallet.Quantity,
	}

	return dto
}
