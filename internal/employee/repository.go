package employee

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e Employee) (int64, error) {
	return 0, nil
}

func (r *Repository) FindByID(id int64) (Employee, error) {
	return Employee{}, nil
}
