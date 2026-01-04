package salary

func CalculateNetSalary(country string, gross float64) float64 {
	if country == "India" {
		return gross * 0.9
	}
	if country == "United States" {
		return gross * 0.88
	}
	return gross
}
