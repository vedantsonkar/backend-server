package routes

import (
	"backend-server/services"
	"backend-server/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	r.POST("/users", func(c *gin.Context) {
		var payload struct {
			UserID string `json:"userId"`
			Email  string `json:"email"`
			Name   string `json:"name"`
		}
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.JSONWithOptionalDebug(c, http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		var dbStart time.Time
		if debug, _ := c.Get("debugEnabled"); debug == true {
			dbStart = time.Now()
		}

		err := services.CreateUser(payload.UserID, payload.Email, payload.Name)

		if debug, _ := c.Get("debugEnabled"); debug == true {
			dbTime := time.Since(dbStart)
			c.Set("dbTime", dbTime)
		}

		if err != nil {
			utils.JSONWithOptionalDebug(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		utils.JSONWithOptionalDebug(c, http.StatusOK, gin.H{"message": "User created"})
	})
}
