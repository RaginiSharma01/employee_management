package router

import (
	"main/handler"
	"net/http"
)

// group employees
func SetupRoutes(employeeHandler *handler.EmployeeHandler) {
	http.HandleFunc("/employees", employeeHandler.CreateEmployee)
	http.HandleFunc("/employees/all", employeeHandler.GetEmployees)
	http.HandleFunc("/employees/get", employeeHandler.GetEmployeeByID)
	http.HandleFunc("/employees/update", employeeHandler.UpdateEmployee)
	http.HandleFunc("/employees/delete", employeeHandler.DeleteEmployee)
	http.HandleFunc("/employees/department", employeeHandler.GetEmployeebyDepartMent)
	http.HandleFunc("/employees/salary", employeeHandler.GetEmployeeFromSalary)
	http.HandleFunc("/employees/department/count", employeeHandler.CountEmployeesByDepartment)
	http.HandleFunc("/employees/recent", employeeHandler.GetRecentEmployees)
	http.HandleFunc("/employees/topSalary", employeeHandler.GetTopPaidEmployees)
}
