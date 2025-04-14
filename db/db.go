package db

import (
	"fmt"
	"gemini_backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	err = database.AutoMigrate(&models.User{})
	if err != nil {
		panic("failed to migrate database: " + err.Error())
	}

	DB = database
	fmt.Println("Database connected and migrated.")
}
