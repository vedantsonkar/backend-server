package controllers

import (
	"errors"
	"net/http"
	"time"

	"backend-server/services"
	"backend-server/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// POST /api/users
func CreateUserHandler(c *gin.Context) {
	var payload struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.JSONWithOptionalDebug(c, http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userID := uuid.New().String() // 🔥 Auto-generate UUID

	var dbStart time.Time
	if debug, _ := c.Get("debugEnabled"); debug == true {
		dbStart = time.Now()
	}

	err := services.CreateUser(userID, payload.Email, payload.Name)

	if debug, _ := c.Get("debugEnabled"); debug == true {
		c.Set("dbTime", time.Since(dbStart))
	}

	if err != nil {
		if errors.Is(err, services.ErrEmailAlreadyExists) {
			utils.JSONWithOptionalDebug(c, http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		utils.JSONWithOptionalDebug(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := utils.GenerateJWT(userID)
	if err != nil {
		utils.JSONWithOptionalDebug(c, http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	utils.JSONWithOptionalDebug(c, http.StatusOK, gin.H{
		"message": "User created successfully",
		"userId":  userID,
		"token":   token,
	})
}

// PUT /api/users
func UpdateUserHandler(c *gin.Context) {
	var payload struct {
		UserID string `json:"userId"`
		Email  string `json:"email"`
		Name   string `json:"name"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		utils.JSONWithOptionalDebug(c, http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var dbStart time.Time
	if debug, _ := c.Get("debugEnabled"); debug == true {
		dbStart = time.Now()
	}

	err := services.UpdateUser(payload.UserID, payload.Email, payload.Name)

	if debug, _ := c.Get("debugEnabled"); debug == true {
		c.Set("dbTime", time.Since(dbStart))
	}

	if err != nil {
		utils.JSONWithOptionalDebug(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.JSONWithOptionalDebug(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DELETE /api/users
func DeleteUserHandler(c *gin.Context) {
	var payload struct {
		UserID string `json:"userId"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil || payload.UserID == "" {
		utils.JSONWithOptionalDebug(c, http.StatusBadRequest, gin.H{"error": "Missing or invalid UserID"})
		return
	}

	var dbStart time.Time
	if debug, _ := c.Get("debugEnabled"); debug == true {
		dbStart = time.Now()
	}

	err := services.DeleteUser(payload.UserID)

	if debug, _ := c.Get("debugEnabled"); debug == true {
		c.Set("dbTime", time.Since(dbStart))
	}

	if err != nil {
		utils.JSONWithOptionalDebug(c, http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	utils.JSONWithOptionalDebug(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
