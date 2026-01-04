package employee

import (
	"testing"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
)

func TestEmployeeRepository_SaveAndFindByID(t *testing.T) {

	// Create in-memory database
	database, err := db.NewInMemoryDB()
	if err != nil {
		t.Fatalf("failed to create database: %v", err)
	}

	// 2. Create employees table
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

	// Create repository
	repo := NewRepository(database)

	// Test data
	employee := Employee{
		FullName: "Amish Jha",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   1000,
	}

	// Save employee
	id, err := repo.Save(employee)
	if err != nil {
		t.Fatalf("unexpected error saving employee: %v", err)
	}

	// Fetch employee
	saved, err := repo.FindByID(id)
	if err != nil {
		t.Fatalf("unexpected error finding employee: %v", err)
	}

	// Assertions
	if saved.FullName != employee.FullName {
		t.Fatalf("expected full name %s, got %s", employee.FullName, saved.FullName)
	}

	if saved.JobTitle != employee.JobTitle {
		t.Fatalf("expected job title %s, got %s", employee.JobTitle, saved.JobTitle)
	}

	if saved.Country != employee.Country {
		t.Fatalf("expected country %s, got %s", employee.Country, saved.Country)
	}

	if saved.Salary != employee.Salary {
		t.Fatalf("expected salary %v, got %v", employee.Salary, saved.Salary)
	}

}
