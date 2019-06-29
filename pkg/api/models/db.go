package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// This is necessary for gorm dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// getEnv get key environment variable if exist otherwise return defalutValue
func getEnv(env string) string {
	value := os.Getenv(env)
	if len(value) == 0 {
		log.Printf("Environment variable %s not found!", env)
	}
	return value
}

// Connect create a new connection with database and return it
func Connect() *gorm.DB {
	dbUser := getEnv("DATABASE_USER")
	dbPass := getEnv("DATABASE_PASSWORD")
	dbHost := getEnv("DATABASE_HOST")
	dbPort := getEnv("DATABASE_PORT")
	dbName := getEnv("DATABASE_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
