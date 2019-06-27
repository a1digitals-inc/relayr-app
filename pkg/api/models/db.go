package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	// This is necessary for gorm dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Connect create a new connection with database and return it
func Connect() *gorm.DB {
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"))
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
