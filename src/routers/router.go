package routes

import (
    "github.com/gin-gonic/gin"
	"tax-travel-tracker/src/handlers"
	"github.com/swaggo/gin-swagger"
    "github.com/swaggo/files" 
)




func SetupRouter() *gin.Engine {
    router := gin.Default()

    // Define your routes
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// API routes
	router.GET("/ping", handlers.PingHandler)
    router.GET("/hello/:name", handlers.HelloHandler)

    return router
}
