package employee

type Employee struct {
	ID       int64   `json:"id,omitempty"`
	FullName string  `json:"full_name"`
	JobTitle string  `json:"job_title"`
	Country  string  `json:"country"`
	Salary   float64 `json:"salary"`
}

func (e Employee) IsValid() bool {
	return hasText(e.FullName) &&
		hasText(e.JobTitle) &&
		hasText(e.Country) &&
		e.Salary > 0
}

func hasText(s string) bool {
	return s != ""
}
