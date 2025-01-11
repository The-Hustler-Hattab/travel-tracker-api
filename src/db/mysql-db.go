package db

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() error {
	// Define your MySQL connection string
	// Format: username:password@tcp(host:port)/dbname
	dsn := os.Getenv("DB_DSN")
	var err error

	// Open the connection
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Test the connection
	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to database: %v", err)
	}
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)

	log.Info("Database connected successfully!")
	return nil
}
