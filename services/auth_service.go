// This package contains the business logic for our application
package services

// Import necessary packages
import (
	"go-gin-auth-api-starter-kit/models"       // Our data models
	"go-gin-auth-api-starter-kit/repositories" // For database operations
	"go-gin-auth-api-starter-kit/utils"        // For helper functions
)

// Register creates a new user account
// user: The user information to register
// Returns: The created user and any error that occurred
func Register(user models.User) (models.User, error) {
	// Hash the user's password for security
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		// If hashing fails, return an empty user and the error
		return models.User{}, err
	}

	// Replace the plain password with the hashed one
	user.Password = hashedPassword

	// Save the user to the database
	return repositories.CreateUser(user)
}

// Login authenticates a user and generates a JWT token
// email: The user's email address
// password: The user's password
// Returns: A JWT token and any error that occurred
func Login(email, password string) (string, error) {
	// Find the user by their email
	user, err := repositories.GetUserByEmail(email)
	if err != nil {
		// If user not found, return empty token and error
		return "", err
	}

	// Check if the provided password matches the stored hash
	if !utils.CheckPasswordHash(password, user.Password) {
		// If password doesn't match, return empty token and error
		return "", err
	}

	// Generate a JWT token for the authenticated user
	return utils.GenerateJWT(user.Username)
}
