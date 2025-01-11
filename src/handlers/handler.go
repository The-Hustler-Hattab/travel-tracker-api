package handlers

import (
    "github.com/gin-gonic/gin"
    "net/http"
	_ "tax-travel-tracker/docs" 

)

// PingHandler responds with a simple JSON message
func PingHandler(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "pong",
    })
}


// PingHandler example
// @Summary      Ping
// @Description  Check API status
// @Tags         health
// @Success      200  {string}  string  "pong"
// @Router       /ping [get]
func HelloHandler(c *gin.Context) {
    name := c.Param("name")
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, " + name + "!",
    })
}