package main

import (
	"fmt"
	"log"
	"main/config"
	"main/db"
	"main/handler"
	"main/repository"
	"main/router"
	"main/services"
	"net/http"
)

func main() {

	cfg := config.LoadConfig()

	database, err := db.ConnectDb(cfg)
	if err != nil {
		log.Fatal(err)
	}

	EmployeeRepo := repository.NewEmployee(database.Pool)
	EmployeeService := services.NewEmployeeService(EmployeeRepo)
	EmployeeHandler := handler.NewEmployeeHandler(EmployeeService)

	router.SetupRoutes(EmployeeHandler)
	fmt.Println("Server is running on port", cfg.ServerPort)
	err = http.ListenAndServe(cfg.ServerPort, nil)
	if err != nil {
		fmt.Printf("error in starting server: %v", err)
	}

}
