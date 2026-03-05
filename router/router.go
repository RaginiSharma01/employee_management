package router

import (
	"main/handler"
	"net/http"
)

func SetupRoutes(employeeHandler *handler.EmployeeHandler) {
	http.HandleFunc("/employees", employeeHandler.CreateEmployee)
}
