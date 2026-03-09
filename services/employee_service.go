package services

import (
	"context"
	"errors"
	"main/models"
	"main/repository"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		Repo: repo,
	}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee models.Employee) error {
	if employee.Name == "" {
		return errors.New("name is required")
	}
	if employee.Email == "" {
		return errors.New("email is required")
	}

	return s.Repo.CreateEmployee(ctx, employee)
}

func (s *EmployeeService) GetEmployees() ([]models.Employee, error) {
	return s.Repo.GetEmployeeData()
}

func (s *EmployeeService) GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {
	return s.Repo.GetEmployeeByID(ctx, id)
}

func (s *EmployeeService) UpdateEmployee(ctx context.Context, employee models.Employee) error {
	if employee.ID == "" {
		return errors.New("id required")
	}
	return s.Repo.UpdateEmployee(ctx, employee)
}

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id int) error {
	return s.Repo.DeleteEmployee(ctx, id)
}
