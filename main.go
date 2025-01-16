package main

import (
	_ "tax-travel-tracker/docs"
	"tax-travel-tracker/src/db"
	"tax-travel-tracker/src/routers"
	"tax-travel-tracker/src/routers/auth"
	"time"
    "sync"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	setupAuthAndDB(log)


	router := routes.SetupRouter()
	router.Run(":5000")

}

func setupAuthAndDB(log *logrus.Logger) {
	measureExecutionTime("setupAuthAndDB", log, func(){

		var wg sync.WaitGroup
		wg.Add(2) // We have two goroutines to wait for
	
		
		go func() {
			defer wg.Done()
			measureExecutionTime("auth.InitVerifier", log, auth.InitVerifier)
		}()
	
		go func() {
			defer wg.Done()
			measureExecutionTime("initializeDatabase", log, func() {
				initializeDatabase(log)
			})
		}()
	
		// Wait for both goroutines to complete
		wg.Wait()
	})
	
	log.Info("Auth and DB setup completed successfully!")
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

func measureExecutionTime(name string, log *logrus.Logger, fn func()) {
    start := time.Now()
    fn()
    elapsed := time.Since(start)
	log.Infof("[+] %s took %s", name, elapsed)
}
