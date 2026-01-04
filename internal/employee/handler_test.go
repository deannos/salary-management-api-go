package employee

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
)

func TestCreateEmployeeHandler(t *testing.T) {
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
	handler := NewHandler(service)

	payload := map[string]interface{}{
		"full_name": "Amish Jha",
		"job_title": "Engineer",
		"country":   "India",
		"salary":    1000,
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handler.CreateEmployee(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
}

func TestGetEmployeeHandler_ShouldReturnEmployee(t *testing.T) {
	// Arrange
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
	handler := NewHandler(service)

	// Create employee via service
	id, _ := service.Create(Employee{
		FullName: "John Doe",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   1000,
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/employees/"+strconv.FormatInt(id, 10),
		nil,
	)
	w := httptest.NewRecorder()

	// Act
	handler.GetEmployee(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var response Employee
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response.ID != id {
		t.Fatalf("expected id %d, got %d", id, response.ID)
	}
}

func TestGetEmployeeSalaryHandler_ShouldReturnSalaryBreakdown(t *testing.T) {
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
	handler := NewHandler(service)

	// Create employee
	id, _ := service.Create(Employee{
		FullName: "John Doe",
		JobTitle: "Engineer",
		Country:  "India",
		Salary:   1000,
	})

	req := httptest.NewRequest(
		http.MethodGet,
		"/employees/"+strconv.FormatInt(id, 10)+"/salary",
		nil,
	)
	w := httptest.NewRecorder()

	// Action
	handler.GetEmployeeSalary(w, req)

	// Assert
	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}

	var response map[string]float64
	if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if response["gross"] != 1000 {
		t.Fatalf("expected gross 1000, got %v", response["gross"])
	}
	if response["net"] != 900 {
		t.Fatalf("expected net 900, got %v", response["net"])
	}
	if response["deduction"] != 100 {
		t.Fatalf("expected deduction 100, got %v", response["deduction"])
	}
}
