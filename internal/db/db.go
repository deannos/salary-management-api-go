package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func NewInMemoryDB() (*sql.DB, error) {
	return sql.Open("sqlite3", ":memory:")
}
