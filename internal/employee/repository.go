package employee

import "database/sql"

const (
	insertEmployeeQuery = `
		INSERT INTO employees (full_name, job_title, country, salary)
		VALUES (?, ?, ?, ?)
	`

	selectEmployeeByIDQuery string = `
		SELECT id, full_name, job_title, country, salary
		FROM employees
		WHERE id = ?
	`
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Save(e Employee) (int64, error) {
	result, err := r.db.Exec(
		insertEmployeeQuery,
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
	row := r.db.QueryRow(selectEmployeeByIDQuery, id)

	var e Employee
	if err := row.Scan(
		&e.ID,
		&e.FullName,
		&e.JobTitle,
		&e.Country,
		&e.Salary,
	); err != nil {
		return Employee{}, err
	}

	return e, nil
}
