package main

import (
"tax-travel-tracker/src/routers"
"github.com/sirupsen/logrus"
"tax-travel-tracker/src/db"
"github.com/joho/godotenv"
_ "tax-travel-tracker/docs" 


) 
// @title           Travel Tracker API
// @version         1.0
// @description     This is API will be used to track travel for users for the purposes of tax savings.
// @contact.name   Mohammed Hattab
// @contact.url    mtahatb.com
// @contact.email  mohammedhattab97@gmail.com
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