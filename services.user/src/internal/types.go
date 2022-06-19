package user

import "github.com/aliaydins/microservice/services.user/src/entity"

type UserDTO struct {
	ID        int    `json:"id""`
	WalletId  int    `json:"wallet_id"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func mapper(user *entity.User) *UserDTO {
	dto := &UserDTO{
		ID:        user.ID,
		WalletId:  user.WalletId,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	return dto
}
