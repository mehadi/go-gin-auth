// This package handles all the controller logic for our application
package controllers

// Import necessary packages
import (
	"go-gin-auth/models"   // Our data models
	"go-gin-auth/services" // Our business logic
	"net/http"             // For HTTP status codes

	"github.com/gin-gonic/gin" // Web framework
)

// Register handles new user registration
// It receives user data and creates a new account
func Register(c *gin.Context) {
	// Create a variable to hold the user data
	var user models.User

	// Try to read the JSON data from the request into our user variable
	if err := c.ShouldBindJSON(&user); err != nil {
		// If there's an error reading the JSON, send a bad request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Try to register the new user using our service
	newUser, err := services.Register(user)

	if err != nil {
		// If registration fails, send an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// If everything went well, send back the created user with a success status
	c.JSON(http.StatusCreated, gin.H{"user": newUser})
}

// Login handles user authentication
// It checks if the provided email and password are correct
func Login(c *gin.Context) {
	// Create a structure to hold login credentials
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Try to read the JSON data from the request
	if err := c.ShouldBindJSON(&credentials); err != nil {
		// If there's an error reading the JSON, send a bad request response
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Try to login using our service
	token, err := services.Login(credentials.Email, credentials.Password)
	if err != nil {
		// If login fails, send an unauthorized response
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// If login is successful, send back the authentication token
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Dashboard handles the post-login page
// It requires a valid JWT token to access
func Dashboard(c *gin.Context) {
	// Get the username from the context (set by middleware)
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Return a welcome message with the username
	c.JSON(http.StatusOK, gin.H{
		"message":  "Welcome to your dashboard!",
		"username": username,
	})
}
