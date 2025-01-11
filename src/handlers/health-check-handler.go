package handlers

import (
	"net/http"
	_ "tax-travel-tracker/docs"

	"github.com/gin-gonic/gin"
)

// HealthCheckHandler godoc
// @Summary Health Check
// @Description Responds with a simple JSON message indicating the service status
// @Tags HealthCheck
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health-check [get]
func HealthCheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Service is up and running!",
	})
}
