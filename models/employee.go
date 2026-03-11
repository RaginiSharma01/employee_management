package models

import "time"

type Employee struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Department  string    `json:"department"`
	Salary      float64   `json:"salary"`
	JoiningDate *time.Time `json:"joining_date,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

// type EmployeeSalaryResponse struct {
// 	ID         string  `json:"id"`
// 	Name       string  `json:"name"`
// 	Email      string  `json:"email"`
// 	Department string  `json:"department"`
// 	Salary     float64 `json:"salary"`
// }
