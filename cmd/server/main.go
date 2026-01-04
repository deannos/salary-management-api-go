package main

import (
	"log"
	"net/http"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
	"github.com/deannos/incubyte-sm-kata-deannos/internal/employee"
	"github.com/deannos/incubyte-sm-kata-deannos/internal/metrics"
)

func main() {
	database, err := db.NewInMemoryDB()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS employees (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			full_name TEXT,
			job_title TEXT,
			country TEXT,
			salary REAL
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	repo := employee.NewRepository(database)
	service := employee.NewService(repo)
	handler := employee.NewHandler(service)

	http.HandleFunc("/employees", handler.CreateEmployee)
	http.HandleFunc("/employees/", handler.EmployeeRoutes)

	metricsHandler := metrics.NewHandler(database)

	http.HandleFunc("/metrics/country/", metricsHandler.GetCountryMetrics)
	http.HandleFunc("/metrics/job-title/", metricsHandler.GetJobTitleMetrics)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
