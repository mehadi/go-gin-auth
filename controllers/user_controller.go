package controllers

import (
	"go-gin-auth-api-starter-kit/config"
	"go-gin-auth-api-starter-kit/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ListUsers returns a list of all users
// This is a protected route that requires authentication
func ListUsers(c *gin.Context) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	// Don't return password hashes in the response
	type UserResponse struct {
		ID        uint   `json:"id"`
		Username  string `json:"username"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
	}

	var response []UserResponse
	for _, user := range users {
		response = append(response, UserResponse{
			ID:        user.ID,
			Username:  user.Username,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, gin.H{"users": response})
}
