package employee

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e Employee) (int64, error) {
	result, err := r.db.Exec(
		`INSERT INTO employees (full_name, job_title, country, salary)
		 VALUES (?, ?, ?, ?)`,
		e.FullName,
		e.JobTitle,
		e.Country,
		e.Salary,
	)
	if err != nil {
		return 0, err
	}

	return result.LastInsertId()
}

func (r *Repository) FindByID(id int64) (Employee, error) {
	row := r.db.QueryRow(
		`SELECT id, full_name, job_title, country, salary
		 FROM employees WHERE id = ?`, id,
	)

	var e Employee
	err := row.Scan(&e.ID, &e.FullName, &e.JobTitle, &e.Country, &e.Salary)
	if err != nil {
		return Employee{}, err
	}

	return e, nil
}
