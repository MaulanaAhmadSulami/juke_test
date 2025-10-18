package employee

import (
	"context"
	"errors"
	"strings"

	employeeEntity "github.com/MaulanaAhmadSulami/juke_test.git/internal/entities/employees"
	"github.com/MaulanaAhmadSulami/juke_test.git/internal/repository/postgres"
)

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *employeeService {
	return &employeeService{
		repo: repo,
	}
}

func (e *employeeService) GetAll(ctx context.Context) ([]employeeEntity.Employee, error) {
	return e.repo.GetAll(ctx)
}

func (e *employeeService) GetById(ctx context.Context, id int64) (*employeeEntity.Employee, error){
	if id <= 0 {
		return nil, errors.New("invalid employee id")
	}

	return e.repo.GetById(ctx, id)
}

func (e *employeeService) Create(ctx context.Context, emp *employeeEntity.Employee) error {
	//vlidaiton
	if strings.TrimSpace(emp.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(emp.Email) == "" {
		return errors.New("email is required")
	}
	if emp.Salary < 0 || emp.Salary == 0 {
		return errors.New("invalid salary")
	}

	emp.Name = strings.TrimSpace(emp.Name)
	emp.Email = strings.ToLower(strings.TrimSpace(emp.Email))
	emp.Position = strings.TrimSpace(emp.Position)

	return e.repo.Create(ctx, emp)
}

func (e *employeeService) Update(ctx context.Context, emp *employeeEntity.Employee) error {
	// Validation
	if emp.ID <= 0 {
		return errors.New("invalid employee ID")
	}
	if strings.TrimSpace(emp.Name) == "" {
		return errors.New("name is required")
	}
	if strings.TrimSpace(emp.Email) == "" {
		return errors.New("email is required")
	}
	if emp.Salary < 0 || emp.Salary == 0{
		return errors.New("invalid salary")
	}

	// Normalize data
	emp.Name = strings.TrimSpace(emp.Name)
	emp.Email = strings.ToLower(strings.TrimSpace(emp.Email))
	emp.Position = strings.TrimSpace(emp.Position)

	return e.repo.Update(ctx, emp)
}

func (e *employeeService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("invalid employee id")
	}


	_, err := e.repo.GetById(ctx, id)
	if err != nil {
		return err
	}
	return e.repo.Delete(ctx, id)
}