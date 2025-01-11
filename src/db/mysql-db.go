package db

import (
	"fmt"
	"log"
	"os"
	"tax-travel-tracker/src/db/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	// Define your MySQL connection string
	dsn := os.Getenv("DB_DSN")
	var err error

	// Open the connection using GORM
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Auto-migrate the schema
	if err := DB.AutoMigrate(&models.TaxTravelTracker{}); err != nil {
		return fmt.Errorf("error during auto-migration: %v", err)
	}

	log.Println("Database connected and schema migrated successfully!")
	return nil
}
