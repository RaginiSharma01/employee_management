package repository

import (
	"context"
	"main/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type EmployeeRepository struct {
	DB *pgxpool.Pool
}

func NewEmployee(pool *pgxpool.Pool) *EmployeeRepository {
	return &EmployeeRepository{
		DB: pool,
	}
}

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, employee models.Employee) error {

	query := `
	INSERT INTO employees_data
	(name, email, department, salary, joining_date, created_at, updated_at)
	VALUES ($1,$2,$3,$4,NOW(),NOW(),NOW())
	`

	_, err := r.DB.Exec(
		ctx,
		query,
		employee.Name,
		employee.Email,
		employee.Department,
		employee.Salary,
	)

	return err
}
//getAll tables (select *from employees_data)

//func (r*EmployeeRepository) GetEmployeeData()
