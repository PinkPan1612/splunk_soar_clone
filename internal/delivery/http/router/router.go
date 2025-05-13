package router

import (
	"net/http"
	"splunk_soar_clone/internal/delivery/http/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/login", authHandler.Login)

	// Protected routes
	auth := r.Group("/api")
	auth.Use(authMiddleware())
	{
		// Add protected routes here
	}

	return r
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		// Validate JWT token here
		// Set user claims in context
		c.Next()
	}
}
