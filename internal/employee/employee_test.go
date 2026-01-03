package employee

import "testing"

func TestEmployee_IsValid_WhenAllFieldsPresent(t *testing.T) {
	e := Employee{
		FullName: "Amish Jha",
		JobTitle: "Software Engineer",
		Country:  "India",
		Salary:   1000,
	}

	if !e.IsValid() {
		t.Fatalf("expected employee to be valid")
	}
}
