package service

import (
	"context"

	employeeEntity "github.com/MaulanaAhmadSulami/juke_test.git/internal/entities/employees"
)

type EmployeesService interface {
	GetAll(context.Context) ([]employeeEntity.Employee, error)
	GetById(context.Context, int64) (*employeeEntity.Employee, error)
	Create(context.Context, *employeeEntity.Employee) error
	Update(context.Context, *employeeEntity.Employee) error
	Delete(context.Context, int64) error
}


type Service struct {
	EmployeesService EmployeesService
}