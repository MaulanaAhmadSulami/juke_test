//Global Error Handling
//Controller advice
package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"
	employeeEntity "github.com/MaulanaAhmadSulami/juke_test.git/internal/entities/employees"
)

var(
	ErrNotFound = errors.New("NOT FOUND")
	QueryTimeoutDuration = time.Second * 5
	ErrNullEmail = errors.New("email cannot be null")
	ErrUniqueViolation = errors.New("an employee wiht this memail already exists")
	ErrNullOrNegSalary = errors.New("salary cannot be null or negative")
)

type Repository struct {
	Employee EmployeeRepository
}

type EmployeeRepository interface {
	GetAll(context.Context) ([]employeeEntity.Employee, error)
	GetById(context.Context, int64) (*employeeEntity.Employee, error)
	Create(context.Context, *employeeEntity.Employee) error
	Update(context.Context, *employeeEntity.Employee) error
	Delete(context.Context, int64) error
}

func WithTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}