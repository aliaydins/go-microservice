package wallet

import "github.com/aliaydins/microservice/service.wallet/src/entity"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetWallet(id int) (*entity.Wallet, error) {
	wallet, err := s.repo.FindById(id)
	if err != nil {
		return nil, ErrWalletNotFound
	}

	return wallet, nil
}
