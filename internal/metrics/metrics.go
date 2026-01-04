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
	row := s.db.QueryRow(`
		SELECT 
			MIN(salary),
			MAX(salary),
			AVG(salary)
		FROM employees
		WHERE country = ?
	`, country)

	var result CountrySalaryMetrics
	err := row.Scan(&result.Min, &result.Max, &result.Avg)
	if err != nil {
		return CountrySalaryMetrics{}, err
	}

	return result, nil
}
