package main

import (
	"log"
	"os"

	"github.com/andrleite/relayr-app/pkg/api"
)

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(env string) string {
	value := os.Getenv(env)
	if len(value) == 0 {
		log.Printf("Environment variable %s not found!", env)
	}
	return value
}

func main() {
	dbUser := getEnv("DATABASE_USER")
	dbPass := getEnv("DATABASE_PASSWORD")
	dbHost := getEnv("DATABASE_HOST")
	dbPort := getEnv("DATABASE_PORT")
	dbName := getEnv("DATABASE_NAME")

	a := api.Api{}
	a.Initialize(dbUser, dbPass, dbHost, dbPort, dbName)
	a.Run()
}
