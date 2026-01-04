package main

import (
	"log"
	"net/http"

	"github.com/deannos/incubyte-sm-kata-deannos/internal/db"
	"github.com/deannos/incubyte-sm-kata-deannos/internal/employee"
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

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
