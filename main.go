package main

import (
"tax-travel-tracker/src/routers"
"github.com/sirupsen/logrus"
"tax-travel-tracker/src/db"
"github.com/joho/godotenv"
_ "tax-travel-tracker/docs" 


) 
// @title           My API
// @version         1.0
// @description     This is a sample API documentation.
// @termsOfService  http://example.com/terms/

// @contact.name   API Support
// @contact.url    http://example.com/support
// @contact.email  support@example.com
func main() {
	log := logrus.New()
	log.SetLevel(logrus.InfoLevel)
	log.Info("Starting the application...")


	loadEnvironmentVariables(log)

	initializeDatabase(log)


	router := routes.SetupRouter()
	router.Run(":8090")



}

func initializeDatabase(log *logrus.Logger) {
	err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
}

func loadEnvironmentVariables(log *logrus.Logger) {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatalf("Error loading .env file: %v", envErr)
	}
	log.Info("Environment variables loaded successfully!")
}