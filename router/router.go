package router

import (
	"main/handler"
	"net/http"
)

func SetupRoutes(employeeHandler *handler.EmployeeHandler) *http.ServeMux {

	mux := http.NewServeMux()

	// Employee Route Group
	mux.HandleFunc("/employees", employeeHandler.CreateEmployee)
	mux.HandleFunc("/employees/all", employeeHandler.GetEmployees)
	mux.HandleFunc("/employees/get", employeeHandler.GetEmployeeByID)
	mux.HandleFunc("/employees/update", employeeHandler.UpdateEmployee)
	mux.HandleFunc("/employees/delete", employeeHandler.DeleteEmployee)

	// Employee Department Routes
	mux.HandleFunc("/employees/department", employeeHandler.GetEmployeebyDepartMent)
	mux.HandleFunc("/employees/department/count", employeeHandler.CountEmployeesByDepartment)

	// Salary Routes
	mux.HandleFunc("/employees/salary", employeeHandler.GetEmployeeFromSalary)
	mux.HandleFunc("/employees/topSalary", employeeHandler.GetTopPaidEmployees)

	// Other Routes
	mux.HandleFunc("/employees/recent", employeeHandler.GetRecentEmployees)

	return mux
}
