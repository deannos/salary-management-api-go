package employee

import (
	"testing"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
)

func TestEmployeeService_Create_ShouldPersistEmployee(t *testing.T) {
	// Arrange: database
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

	repo := NewRepository(database)
	service := NewService(repo)

	// Act
	id, err := service.Create(Employee{
		FullName: "Sahil Khan",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   2000,
	})

	// Assert
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if id == 0 {
		t.Fatalf("expected non-zero employee id")
	}
}

func TestEmployeeService_Create_ShouldFailForInvalidEmployee(t *testing.T) {
	database, _ := db.NewInMemoryDB()

	_, _ = database.Exec(`
		CREATE TABLE employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT,
			job_title TEXT,
			country TEXT,
			salary REAL
		)
	`)

	repo := NewRepository(database)
	service := NewService(repo)

	_, err := service.Create(Employee{})
	if err == nil {
		t.Fatalf("expected error for invalid employee")
	}
}

func TestEmployeeService_GetByID_ShouldReturnEmployee(t *testing.T) {
	database, _ := db.NewInMemoryDB()

	_, _ = database.Exec(`
		CREATE TABLE employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT,
			job_title TEXT,
			country TEXT,
			salary REAL
		)
	`)

	repo := NewRepository(database)
	service := NewService(repo)

	id, _ := service.Create(Employee{
		FullName: "John Smith",
		JobTitle: "Manager",
		Country:  "US",
		Salary:   3000,
	})

	e, err := service.GetByID(id)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if e.FullName != "John Smith" {
		t.Fatalf("unexpected employee returned")
	}
}
