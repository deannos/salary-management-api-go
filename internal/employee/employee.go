package employee

type Employee struct {
	FullName string
	JobTitle string
	Country  string
	Salary   float64
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
