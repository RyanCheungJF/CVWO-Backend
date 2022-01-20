package database

import (
	"os"
	"fmt"
	
	"github.com/RyanCheungJF/CVWO-Backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Defining global var for connection (referencing DB connection)
var DB *gorm.DB

func Connect() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("HOST")
	database := os.Getenv("DATABASE")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", username, password, host, database)
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Could not connect to the database.")
	}

	DB = connection

	// Database Migration for both tables
	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Task{})
}
