package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type ConfigStruct struct {
	DBIP       string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string
	ServerPort string
}

func LoadConfig() *ConfigStruct {
	err := godotenv.Load("config/.env")

	if err != nil {
		fmt.Printf("error in loading to .env file")
	}

	return &ConfigStruct{
		DBIP:       os.Getenv("DB_IP"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		ServerPort: os.Getenv("SERVER_PORT"),
	}

}
