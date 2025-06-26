package routes

import (
	"backend-server/controllers"
	"backend-server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Open route with rate limiting
	r.GET("/api/latency", middleware.RateLimitMiddleware(), controllers.LatencyHandler)

	// Secure group (JWT auth)
	authGroup := r.Group("/api/secure")
	authGroup.Use(middleware.RateLimitMiddleware(), middleware.AuthMiddleware())
	{
		authGroup.GET("/protected", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "You are authenticated"})
		})
	}
}
