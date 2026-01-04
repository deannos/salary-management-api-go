package metrics

import (
	"testing"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
	"github.com/deannos/incubyte-sm-kata-deannos/internal/employee"
)

func TestSalaryMetrics_ByCountry(t *testing.T) {
	database, err := db.NewInMemoryDB()
	if err != nil {
		t.Fatalf("failed to create db: %v", err)
	}

	_, err = database.Exec(`
		CREATE TABLE employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT,
			job_title TEXT,
			country TEXT,
			salary REAL
		)
	`)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}

	repo := employee.NewRepository(database)
	service := employee.NewService(repo)

	_, _ = service.Create(employee.Employee{
		FullName: "A",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   1000,
	})
	_, _ = service.Create(employee.Employee{
		FullName: "B",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   2000,
	})
	_, _ = service.Create(employee.Employee{
		FullName: "C",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   3000,
	})

	metrics := NewService(database)

	result, err := metrics.ByCountry("India")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Min != 1000 {
		t.Fatalf("expected min 1000, got %v", result.Min)
	}
	if result.Max != 3000 {
		t.Fatalf("expected max 3000, got %v", result.Max)
	}
	if result.Avg != 2000 {
		t.Fatalf("expected avg 2000, got %v", result.Avg)
	}
}
