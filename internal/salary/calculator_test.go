package salary

import "testing"

func TestCalculateNetSalary_ForIndia(t *testing.T) {
	net := CalculateNetSalary("India", 1000)

	if net != 900 {
		t.Fatalf("expected net salary 900, got %v", net)
	}
}

func TestCalculateNetSalary_ForUnitedStates(t *testing.T) {
	net := CalculateNetSalary("United States", 1000)

	if net != 880 {
		t.Fatalf("expected net salary 880, got %v", net)
	}
}
