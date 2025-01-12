package routes

import (
	"tax-travel-tracker/src/handlers"
	"tax-travel-tracker/src/handlers/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

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



	return router
}
