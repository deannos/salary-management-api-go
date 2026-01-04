package metrics

import "database/sql"

type Service struct {
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

type CountrySalaryMetrics struct {
	Min float64
	Max float64
	Avg float64
}

func (s *Service) ByCountry(country string) (CountrySalaryMetrics, error) {
	return CountrySalaryMetrics{}, nil
}
