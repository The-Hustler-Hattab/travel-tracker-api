package routes

import (
	"tax-travel-tracker/src/handlers"
	"tax-travel-tracker/src/handlers/middleware"
	// "time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()



	    // Option A: Use the default CORS settings (allows *all* origins, methods, headers)
    router.Use(cors.Default())

    // Option B: Configure CORS for specific origins, methods, etc.
    // router.Use(cors.New(cors.Config{
    //     AllowOrigins:     []string{"http://localhost:4200"}, // or "*" to allow all
    //     AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
    //     AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
    //     ExposeHeaders:    []string{"Content-Length"},
    //     AllowCredentials: true,
    //     MaxAge:           12 * time.Hour,
    // }))


	// Define your routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	router.GET("/health-check", handlers.HealthCheckHandler)


    	// Travel Tracker API routes
	travelGroup := router.Group("/travel-tracker")
	travelGroup.Use(middleware.AuthMiddleware())
	{
		travelGroup.POST("", handlers.CreateTravelRecordHandler)
		travelGroup.GET("/:id", handlers.GetTravelRecordHandler)
		travelGroup.PUT("/:id", handlers.UpdateTravelRecordHandler)
		travelGroup.DELETE("/:id", handlers.DeleteTravelRecordHandler)
		travelGroup.GET("", handlers.GetAllTravelRecordsHandler)

	}

	milageGroup := router.Group("/mileage")
	milageGroup.Use(middleware.AuthMiddleware())
	{
		milageGroup.GET("/:id", handlers.GetMilageRecordHandler)
		milageGroup.GET("", handlers.GetAllMileageRecordsHandler)

	}

	apiGroup := router.Group("/api-key")
	apiGroup.Use(middleware.AuthMiddleware())
	{
		apiGroup.GET("", handlers.ApiKeyHandler)
	}

	userGroup := router.Group("/users")
	{
		userGroup.POST("", handlers.CreateUserHandler)
	}



	return router
}
