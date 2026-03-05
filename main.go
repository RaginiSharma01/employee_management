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

	fmt.Println("server is listening on localhost:9001")

	log.Fatal(http.ListenAndServe(":9001", nil))
}
