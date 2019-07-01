package models

import (
	"fmt"

	"github.com/jinzhu/gorm"

	// This is necessary for gorm dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

//Database struct to connect to gorm
type Database struct {
	*gorm.DB
}

//New return new connection pool
func New(dbUser, dbPass, dbHost, dbPort, dbName string) (*Database, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName)
	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "unable to connect to database")
	}
	return &Database{db}, nil
}
