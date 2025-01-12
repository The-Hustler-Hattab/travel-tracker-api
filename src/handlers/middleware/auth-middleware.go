package middleware

import (
	"net/http"
	"strings"
	"tax-travel-tracker/src/routers/auth"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
            return
        }

        token := strings.TrimPrefix(authHeader, "Bearer ")
        if token == authHeader {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Bearer token missing"})
            return
        }

        _, err := auth.JwtVerifier.VerifyAccessToken(token)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Next()
    }
}
