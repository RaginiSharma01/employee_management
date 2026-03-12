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

func (r *EmployeeRepository) CreateEmployee(ctx context.Context, employee models.Employee) (string, error) {

	query := `
	INSERT INTO employees_data
	(name, email, department, salary, joining_date, created_at, updated_at)
	VALUES ($1,$2,$3,$4,NOW(),NOW(),NOW())
	RETURNING id
	`

	var id string

	err := r.DB.QueryRow(
		ctx,
		query,
		employee.Name,
		employee.Email,
		employee.Department,
		employee.Salary,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

//getAll tables (select *from employees_data)

func (r *EmployeeRepository) GetEmployeeData(ctx context.Context, limit, offset int) ([]models.Employee, error) {

	query := `
	SELECT id, name, email, department, salary, joining_date
	FROM employees_data
	ORDER BY created_at DESC
	LIMIT $1 OFFSET $2
	`

	rows, err := r.DB.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Department,
			&emp.Salary,
			&emp.JoiningDate,
		)
		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

func (r *EmployeeRepository) GetEmployeeByID(ctx context.Context, id int) (*models.Employee, error) {

	query := `SELECT id, name, email, department, salary, joining_date, created_at, updated_at
			  FROM employees_data WHERE id=$1`

	var employee models.Employee

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Department,
		&employee.Salary,
		&employee.JoiningDate,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (r *EmployeeRepository) UpdateEmployee(ctx context.Context, employee models.Employee) error {

	query := `
	UPDATE employees_data
	SET name=$1, email=$2, department=$3, salary=$4, updated_at=NOW()
	WHERE id=$5
	`

	_, err := r.DB.Exec(ctx, query,
		employee.Name,
		employee.Email,
		employee.Department,
		employee.Salary,
		employee.ID,
	)

	return err
}

func (r *EmployeeRepository) DeleteEmployee(ctx context.Context, id string) error {

	query := `DELETE FROM employees_data WHERE id=$1`

	_, err := r.DB.Exec(ctx, query, id)

	return err
}

//get employees from department/salaray

func (r *EmployeeRepository) GetEmployeebyDepartMent(dept string) ([]models.Employee, error) {
	query := `SELECT*FROM employees_data WHERE department =$1`
	rows, err := r.DB.Query(context.Background(), query, dept)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee
	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Department,
			&emp.Salary,
			&emp.JoiningDate,
			&emp.CreatedAt,
			&emp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

// call employees based on the salaray

func (r *EmployeeRepository) GetEmployeeFromSalary(amount float64) ([]models.Employee, error) {
	query := `SELECT id, name, email, department, salary
	FROM employees_data
	WHERE salary >= $1
	`

	rows, err := r.DB.Query(context.Background(), query, amount)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Department,
			&emp.Salary,
		)
		if err != nil {
			return nil, err
		}
		employees = append(employees, emp)
	}
	return employees, nil
}

func (r *EmployeeRepository) CountEmployeesByDepartment() (map[string]int, error) {

	query := `
	SELECT department, COUNT(*)
	FROM employees_data
	GROUP BY department
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]int)

	for rows.Next() {
		var dept string
		var count int

		err := rows.Scan(&dept, &count)
		if err != nil {
			return nil, err
		}

		result[dept] = count
	}

	return result, nil
}

func (r *EmployeeRepository) GetRecentEmployees() ([]models.Employee, error) {

	query := `
	SELECT id, name, email, department, salary, joining_date, created_at, updated_at
	FROM employees_data
	WHERE joining_date >= NOW() - INTERVAL '2 days'
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee

		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Department,
			&emp.Salary,
			&emp.JoiningDate,
			&emp.CreatedAt,
			&emp.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}

//get top 5 employees;

func (r *EmployeeRepository) GetTopPaidEmployees() ([]models.Employee, error) {

	query := `
	SELECT id, name, email, department, salary, joining_date, created_at, updated_at
	FROM employees_data
	ORDER BY salary DESC
	LIMIT 5
	`

	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employees []models.Employee

	for rows.Next() {
		var emp models.Employee

		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Email,
			&emp.Department,
			&emp.Salary,
			&emp.JoiningDate,
			&emp.CreatedAt,
			&emp.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		employees = append(employees, emp)
	}

	return employees, nil
}
