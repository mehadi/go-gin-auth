// This package contains all our data models
package models

// Import GORM for database operations
import "gorm.io/gorm"

// User represents a person who can use our application
// It includes basic information like username, email, and password
type User struct {
	// gorm.Model adds these fields automatically:
	// - ID (unique identifier)
	// - CreatedAt (when the user was created)
	// - UpdatedAt (when the user was last updated)
	// - DeletedAt (when the user was deleted, if they were)
	gorm.Model

	// Username is the name the user will use to log in
	// It must be unique (no two users can have the same username)
	// It cannot be empty
	Username string `gorm:"unique;not null" json:"username"`

	// Email is the user's email address
	// It must be unique (no two users can have the same email)
	// It cannot be empty
	Email    string `gorm:"unique;not null" json:"email"`

	// Password is the user's secret password
	// It cannot be empty
	// Note: In a real application, this should be hashed before storing
	Password string `gorm:"not null" json:"password"`
}
