package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v3/log"
)

type EmployeeHandler struct {
	Service *services.EmployeeService
}

func NewEmployeeHandler(service *services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		Service: service,
	}
}

func writeJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSONResponse(w, status, map[string]string{
		"error": message,
	})
}

func (h *EmployeeHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	defer r.Body.Close()

	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json payload")
		return
	}

	id, err := h.Service.CreateEmployee(r.Context(), employee)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]string{
		"message": "employee created successfully",
		"id":      id,
	})
}

func (h *EmployeeHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	employees, err := h.Service.GetEmployees()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var response []models.EmployeeResPonse

	for _, emp := range employees {
		var joiningDate string

		if emp.JoiningDate != nil {
			joiningDate = emp.JoiningDate.Format("2006-01-02")
		}

		response = append(response, models.EmployeeResPonse{
			ID:          emp.ID,
			Name:        emp.Name,
			Email:       emp.Email,
			Department:  emp.Department,
			Salary:      emp.Salary,
			JoiningDate: joiningDate,
		})
	}

	writeJSONResponse(w, http.StatusOK, response)
}
func (h *EmployeeHandler) GetEmployeeByID(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		writeJSONError(w, http.StatusBadRequest, "id is required")
		return
	}

	id, _ := strconv.Atoi(idStr)

	employee, err := h.Service.GetEmployeeByID(r.Context(), id)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSONResponse(w, http.StatusOK, employee)
}

func (h *EmployeeHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var employee models.Employee

	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	err = h.Service.UpdateEmployee(r.Context(), employee)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]string{
		"message": "employee updated successfully",
	})
}

func (h *EmployeeHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Query().Get("id")

	if idStr == "" {
		writeJSONError(w, http.StatusBadRequest, "id required")
		return
	}

	err := h.Service.DeleteEmployee(r.Context(), idStr)
	if err != nil {
		log.Errorf("db error", err.Error())
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSONResponse(w, http.StatusOK, map[string]string{
		"message": "employee deleted successfully",
	})
}

// employ by department handler

//method ko call karke check karo

func (h *EmployeeHandler) GetEmployeebyDepartMent(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	dept := r.URL.Query().Get("dept")

	employees, err := h.Service.GetEmployeebyDepartMent(dept)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, http.StatusOK, employees)
}

//handler to manage the get employee from there salary api

func (h *EmployeeHandler) GetEmployeeFromSalary(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	amountStr := r.URL.Query().Get("amount")

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		http.Error(w, "invalid salary", http.StatusBadRequest)
		return
	}

	employees, err := h.Service.GetEmployeeFromSalary(amount)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSONResponse(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) CountEmployeesByDepartment(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	result, err := h.Service.CountEmployeesByDepartment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, http.StatusOK, result)
}

func (h *EmployeeHandler) GetRecentEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	employees, err := h.Service.GetRecentEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, http.StatusOK, employees)
}

func (h *EmployeeHandler) GetTopPaidEmployees(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	employees, err := h.Service.GetTopPaidEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	writeJSONResponse(w, http.StatusOK, employees)
}
