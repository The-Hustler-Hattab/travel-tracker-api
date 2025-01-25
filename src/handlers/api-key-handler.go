package handlers

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

// ApiKeyController godoc
// @Summary Get API Key
// @Description Responds with a simple JSON message that has the api key
// @Tags ApiKey
// @Produce json
// @Success 200 {object} map[string]string
// @Router /api-key [get]
func ApiKeyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"apiKey": os.Getenv("GOOGLE_MAP_API_KEY"),
	})
}
