package database

import (
	"go-api/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("database/users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos")
	}

	// Migrar el modelo de usuario para crear la tabla
	DB.AutoMigrate(&models.User{})
}
