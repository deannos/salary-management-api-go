package employee

type Employee struct {
	FullName string
	JobTitle string
	Country  string
	Salary   float64
}

func (e Employee) IsValid() bool {
	return true
}
