package metrics

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
	"github.com/deannos/incubyte-sm-kata-deannos/internal/employee"
)

func TestCountryMetricsHandler(t *testing.T) {
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
		Salary:   3000,
	})

	handler := NewHandler(database)

	req := httptest.NewRequest(http.MethodGet, "/metrics/country/India", nil)
	w := httptest.NewRecorder()

	handler.GetCountryMetrics(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	var resp map[string]float64
	if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response")
	}

	if resp["min"] != 1000 {
		t.Fatalf("expected min 1000, got %v", resp["min"])
	}
	if resp["max"] != 3000 {
		t.Fatalf("expected max 3000, got %v", resp["max"])
	}
	if resp["avg"] != 2000 {
		t.Fatalf("expected avg 2000, got %v", resp["avg"])
	}
}
