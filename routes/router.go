package routes

import (
	"backend-server/controllers"
	"backend-server/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public endpoints
	public := r.Group("/api")
	public.Use(middleware.RateLimitMiddleware(), middleware.DebugTiming())
	{
		public.GET("/latency", controllers.LatencyHandler)
		public.POST("/users", controllers.CreateUserHandler) // ✅ no JWT needed
	}

	// Secure endpoints
	secure := r.Group("/api")
	secure.Use(middleware.RateLimitMiddleware(), middleware.DebugTiming(), middleware.AuthMiddleware())
	{
		secure.PUT("/users", controllers.UpdateUserHandler)
		secure.DELETE("/users", controllers.DeleteUserHandler)
	}
}
