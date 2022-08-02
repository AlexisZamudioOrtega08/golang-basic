package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string = "host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"
var DB *gorm.DB
var err error

func DBConnection() {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB Connection Successful")
	}
}
