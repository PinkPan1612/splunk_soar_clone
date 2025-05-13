package router

import (
	"splunk_soar_clone/internal/delivery/http/handler"
	"splunk_soar_clone/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(authHandler *handler.AuthHandler, jwtKey []byte) *gin.Engine {
	r := gin.Default()

	// Public routes
	r.POST("/login", authHandler.Login)

	// Protected routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(jwtKey))
	{
		// Admin only routes
		admin := api.Group("/admin")
		admin.Use(middleware.RoleMiddleware("1")) // Admin role
		{
			// Add admin routes here
		}

		// User routes
		users := api.Group("/users")
		users.Use(middleware.RoleMiddleware("1", "2")) // Admin and regular users
		{
			// Add user routes here
		}
	}

	return r
}
