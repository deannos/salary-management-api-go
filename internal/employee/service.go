package employee

import "errors"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(e Employee) (int64, error) {
	if !e.IsValid() {
		return 0, errors.New("invalid employee")
	}
	return s.repo.Save(e)
}
