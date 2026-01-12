package employee

type Employee struct {
	ID       int64   `json:"id,omitempty"`
	FullName string  `json:"full_name"`
	JobTitle string  `json:"job_title"`
	Country  string  `json:"country"`
	Salary   float64 `json:"salary"`
}

func (e Employee) IsValid() bool {
	if !hasText(e.FullName) ||
		!hasText(e.JobTitle) ||
		!hasText(e.Country) ||
		e.Salary <= 0 {
		return false
	}

	if e.Country == "India" && e.Salary >= 100000 {
		return false
	}
	return true
}

func hasText(s string) bool {
	return s != ""
}
