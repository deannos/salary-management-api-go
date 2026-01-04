package salary

func CalculateNetSalary(country string, gross float64) float64 {
	deductionRate := deductionRateFor(country)
	return gross * (1 - deductionRate)
}

func deductionRateFor(country string) float64 {
	switch country {
	case "India":
		return 0.10
	case "United States":
		return 0.12
	default:
		return 0.0
	}
}
