package main

import (
	"os"

	"backend-server/config"
	"backend-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadAWSConfig()
	
	r := gin.Default()
	routes.SetupRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}
