package handler

import (
	"encoding/json"
	"net/http"

	"main/models"
	"main/services"
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

	if employee.Name == "" || employee.Email == "" {
		writeJSONError(w, http.StatusBadRequest, "name and email required")
		return
	}

	err = h.Service.CreateEmployee(r.Context(), employee)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	writeJSONResponse(w, http.StatusCreated, map[string]string{
		"message": "employee created successfully",
	})
}
