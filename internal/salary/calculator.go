package salary

import "math"

func CalculateNetSalary(country string, gross float64) float64 {
	deductionRate := deductionRateFor(country)
	net := gross * (1 - deductionRate)
	return roundToTwoDecimals(net)
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

func roundToTwoDecimals(value float64) float64 {
	return math.Round(value*100) / 100
}
