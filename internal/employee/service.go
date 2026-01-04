package employee

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(e Employee) (int64, error) {
	return s.repo.Save(e)
}
