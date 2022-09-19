package history

import "github.com/aliaydins/microservice/service.history/src/entity"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Create(history *entity.History) error {
	return s.repo.Create(history)
}

func (s *Service) GetListById(userId int) ([]HistoryDTO, error) {
	h, err := s.repo.GetListById(userId)
	if err != nil {
		return nil, ErrNotFoundAnyData
	}

	hDto := make([]HistoryDTO, 0)

	for _, e := range h {
		hDto = append(hDto, mapper(e))
	}

	return hDto, nil
}
