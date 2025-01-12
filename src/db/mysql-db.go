package db

import (
	"fmt"
	"log"
	"os"
	"time"
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

	
	// Retrieve the underlying sql.DB to configure the connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error retrieving database instance: %v", err)
	}

	// Set the maximum number of idle connections in the pool
	sqlDB.SetMaxIdleConns(5)

	// Set the maximum number of open connections to the database
	sqlDB.SetMaxOpenConns(10)

	// Set the maximum lifetime of a connection
	sqlDB.SetConnMaxLifetime(time.Minute * 5)

	log.Println("Database connected successfully!")
	return nil
}
