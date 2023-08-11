package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=manuel password=mysecretpassword dbname=gorm port=5432 sslmode=disable"
var DB *gorm.DB

func DBConnect() {
	var errorMessage error
	DB, errorMessage = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if errorMessage != nil {
		log.Fatal(errorMessage)
	} else {
		log.Println("Connected to database")
	}
}
