package employee

import (
	"context"
	"database/sql"
	"errors"

	employeeEntity "github.com/MaulanaAhmadSulami/juke_test.git/internal/entities/employees"
	repository "github.com/MaulanaAhmadSulami/juke_test.git/internal/repository/postgres"
)

func NewEmployeeStore(db *sql.DB) *employeeStore {
	return &employeeStore{
		DB: db,
	}
}

type employeeStore struct {
	DB *sql.DB
}

func (e *employeeStore) GetAll(ctx context.Context) ([]employeeEntity.Employee, error) {
	query := `SELECT id, name, email, position, salary, created_at FROM employees`
	
	rows, err := e.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var employees []employeeEntity.Employee
	for rows.Next() {
		var emp employeeEntity.Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Email, &emp.Position, &emp.Salary, &emp.CreatedAt)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	
	return employees, rows.Err()
}

func(e *employeeStore) GetById(ctx context.Context, empid int64) (*employeeEntity.Employee, error) {
	query := `
		SELECT id, name, email, position, salary, created_at
		FROM employees
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, repository.QueryTimeoutDuration)
	defer cancel()


	var emp employeeEntity.Employee
	err := e.DB.QueryRowContext(
		ctx,
		query,
		empid,
	).Scan(
		&emp.ID,
		&emp.Name,
		&emp.Email,
		&emp.Position,
		&emp.Salary,
		&emp.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, repository.ErrNotFound
		default:
			return nil, err
		}
	}

	return &emp, nil

}

func(e *employeeStore) Create(ctx context.Context, emp *employeeEntity.Employee) error {
	query := `
		INSERT INTO employees (name, email, position, salary)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at
	`

	ctx, cancel := context.WithTimeout(ctx, repository.QueryTimeoutDuration)
	defer cancel()

	err := e.DB.QueryRowContext(
		ctx,
		query,
		emp.Name,
		emp.Email,
		emp.Position,
		emp.Salary,
	).Scan(&emp.ID, &emp.CreatedAt)

	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "employees_email_key"`:
			return repository.ErrUniqueViolation
		case err.Error() == `pq: email cannot be null`:
			return repository.ErrNullEmail
		case err.Error() == `pq: salary canot be null or neagative`:
			return repository.ErrNullOrNegSalary
		default:
			return err
		}
	}

	return nil
}

func(e *employeeStore) Update(ctx context.Context, emp *employeeEntity.Employee) error {
	query := `
		UPDATE employees SET
		name = $1,
		email = $2,
		position = $3,
		salary = $4
		WHERE id = $5
	`

	ctx, cancel := context.WithTimeout(ctx, repository.QueryTimeoutDuration)
	defer cancel()

	_, err := e.DB.ExecContext(
		ctx, 
		query, 
		emp.Name,
		emp.Email,
		emp.Position,
		emp.Salary,
		emp.ID);

	if err != nil {
		switch {
			case err.Error() == `pq: duplicate key value violates unique constraint "employees_email_key"`:
				return repository.ErrUniqueViolation
			case err.Error() == `pq: salary canot be null or neagative`:
				return repository.ErrNullOrNegSalary
			default:
				return err
		}
	}

	return nil
}

func(e *employeeStore) Delete(ctx context.Context, empId int64) error {
	query := `
		DELETE FROM employees WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, repository.QueryTimeoutDuration)
	defer cancel()

	result, err := e.DB.ExecContext(ctx, query, empId)
	if err != nil{
		return err
	}

	isAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if isAffected == 0 {
		return repository.ErrNotFound
	}

	return nil
}