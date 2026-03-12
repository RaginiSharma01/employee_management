package models

import "time"

type Employee struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Email       string     `json:"email"`
	Department  string     `json:"department"`
	Salary      float64    `json:"salary"`
	JoiningDate *time.Time `json:"joiningDate,omitempty"`
	CreatedAt   *time.Time `json:"createdAt,omitempty"`
	UpdatedAt   *time.Time `json:"updatedAt,omitempty"`
}

// type EmployeeSalaryResponse struct {
// 	ID         string  `json:"id"`
// 	Name       string  `json:"name"`
// 	Email      string  `json:"email"`
// 	Department string  `json:"department"`
// 	Salary     float64 `json:"salary"`
// }

type EmployeeResPonse struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Email       string  `json:"email"`
	Department  string  `json:"department"`
	Salary      float64 `json:"salary"`
	JoiningDate string  `json:"joiningDate,omitempty"`
}
