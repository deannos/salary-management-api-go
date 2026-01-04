package employee

import (
	"testing"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
)

func TestEmployeeRepository_SaveAndFindByID(t *testing.T) {
	database, err := db.NewInMemoryDB()
	if err != nil {
		t.Fatalf("failed to create database: %v", err)
	}

	repo := NewRepository(database)

	employee := Employee{
		FullName: "Amish Jha",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   1000,
	}

	id, err := repo.Save(employee)
	if err != nil {
		t.Fatalf("unexpected error saving employee: %v", err)
	}

	saved, err := repo.FindByID(id)
	if err != nil {
		t.Fatalf("unexpected error finding employee: %v", err)
	}

	if saved.FullName != employee.FullName {
		t.Fatalf("expected %s, got %s", employee.FullName, saved.FullName)
	}

}
