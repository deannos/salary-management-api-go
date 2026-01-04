package metrics

import "database/sql"

const (
	countryMetricsQuery = `
		SELECT 
			MIN(salary),
			MAX(salary),
			AVG(salary)
		FROM employees
		WHERE country = ?
	`

	jobTitleAverageQuery = `
		SELECT AVG(salary)
		FROM employees
		WHERE job_title = ?
	`
)

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
	row := s.db.QueryRow(countryMetricsQuery, country)

	var result CountrySalaryMetrics
	if err := row.Scan(&result.Min, &result.Max, &result.Avg); err != nil {
		return CountrySalaryMetrics{}, err
	}

	return result, nil
}

func (s *Service) AverageByJobTitle(title string) (float64, error) {
	row := s.db.QueryRow(jobTitleAverageQuery, title)

	var avg float64
	if err := row.Scan(&avg); err != nil {
		return 0, err
	}

	return avg, nil
}
