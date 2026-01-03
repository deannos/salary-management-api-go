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

func TestEmployee_IsInvalid_WhenFullNameMissing(t *testing.T) {
	e := Employee{
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   1000,
	}

	if e.IsValid() {
		t.Fatalf("expected employee to be invalid when full name is missing")
	}
}

func TestEmployee_IsInvalid_WhenJobTitleMissing(t *testing.T) {
	e := Employee{
		FullName: "Amish Jha",
		Country:  "India",
		Salary:   1000,
	}

	if e.IsValid() {
		t.Fatalf("expected employee to be invalid when job title is missing")
	}
}
