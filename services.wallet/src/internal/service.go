package wallet

import (
	"fmt"
	"github.com/aliaydins/microservice/service.wallet/src/entity"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetWallet(id int) (*WalletDTO, error) {
	wallet, err := s.repo.FindById(id)
	if err != nil {
		return nil, ErrWalletNotFound
	}

	walletDto := mapper(wallet)

	return walletDto, nil
}

func (s *Service) CreateWallet(userID int) error {
	newWallet := entity.Wallet{
		UserId: userID,
		USD:    30000,
		BTC:    5,
	}

	err := s.repo.Create(&newWallet)
	if err != nil {
		return ErrWalletNotCreated
	}

	return nil
}

func (s *Service) UpdateWallet(wallet *entity.Wallet, id int) (*WalletDTO, error) {

	err := s.repo.UpdateByUserID(wallet, id)
	if err != nil {
		return nil, ErrWalletNotUpdated
	}
	fmt.Println(wallet)
	return mapper(wallet), nil
}
