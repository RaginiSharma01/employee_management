package repository

import (
	"context"
	"fmt"
	"log"
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

func (r *EmployeeRepository) GetEmployeeData() ([]models.Employee, error) {
	query := "SELECT id , name,email,department,salary,joining_date,created_at FROM employees_data"

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() 


	var employees []models.Employee

	for rows.Next() {
		var employee models.Employee

		rows.Scan(&employee)
		employees = append(employees, employee)
	}

	fmt.Println(employees)
	return employees , nil
}
