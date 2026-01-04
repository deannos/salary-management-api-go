package employee

import "errors"

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func validateEmployee(e Employee) error {
	if !e.IsValid() {
		return errors.New("invalid employee")
	}
	return nil
}

func (s *Service) Create(e Employee) (int64, error) {
	if err := validateEmployee(e); err != nil {
		return 0, err
	}
	return s.repo.Save(e)
}

func (s *Service) GetByID(id int64) (Employee, error) {
	return s.repo.FindByID(id)
}
