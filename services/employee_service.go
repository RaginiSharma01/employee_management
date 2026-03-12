package services

import (
	"context"
	"errors"
	"main/models"
	"main/repository"
	"strings"
)

type EmployeeService struct {
	Repo *repository.EmployeeRepository
}

func NewEmployeeService(repo *repository.EmployeeRepository) *EmployeeService {
	return &EmployeeService{
		Repo: repo,
	}
}

func (s *EmployeeService) CreateEmployee(ctx context.Context, employee models.Employee) (string, error) {

	if employee.Name == "" {
		return "", errors.New("name is required")
	}

	if employee.Email == "" {
		return "", errors.New("email is required")
	}

	employee.Email = strings.ToLower(employee.Email)

	return s.Repo.CreateEmployee(ctx, employee)
}

func (s *EmployeeService) GetEmployees(ctx context.Context, limit, offset int) ([]models.Employee, error) {
	return s.Repo.GetEmployeeData(ctx, limit, offset)
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

func (s *EmployeeService) DeleteEmployee(ctx context.Context, id string) error {
	return s.Repo.DeleteEmployee(ctx, id)
}

func (s *EmployeeService) GetEmployeebyDepartMent(dept string) ([]models.Employee, error) {
	return s.Repo.GetEmployeebyDepartMent(dept)
}

func (s *EmployeeService) GetEmployeeFromSalary(amount float64) ([]models.Employee, error) {
	return s.Repo.GetEmployeeFromSalary(amount)
}

func (s *EmployeeService) CountEmployeesByDepartment() (map[string]int, error) {
	return s.Repo.CountEmployeesByDepartment()
}

func (s *EmployeeService) GetRecentEmployees() ([]models.Employee, error) {
	return s.Repo.GetRecentEmployees()
}

func (s *EmployeeService) GetTopPaidEmployees() ([]models.Employee, error) {
	return s.Repo.GetTopPaidEmployees()
}
