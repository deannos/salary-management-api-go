package employee

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
