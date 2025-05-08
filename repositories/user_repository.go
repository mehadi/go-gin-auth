// This package handles all database operations
package repositories

// Import necessary packages
import (
	"go-gin-auth-api-starter-kit/config" // Database configuration
	"go-gin-auth-api-starter-kit/models" // User model
)

// CreateUser saves a new user to the database
// user: The user information to save
// Returns: The saved user and any error that occurred
func CreateUser(user models.User) (models.User, error) {
	// Use GORM to create a new record in the users table
	err := config.DB.Create(&user).Error
	return user, err
}

// GetUserByEmail finds a user by their email address
// email: The email address to search for
// Returns: The found user and any error that occurred
func GetUserByEmail(email string) (models.User, error) {
	// Create a variable to hold the user
	var user models.User

	// Use GORM to find the first user with matching email
	err := config.DB.Where("email = ?", email).First(&user).Error
	return user, err
}
