// This is the main package where our application starts
package main

// Import necessary packages
import (
	"go-gin-auth/config" // Our database configuration
	"go-gin-auth/models" // Our data models (like User)
	"go-gin-auth/routes" // Our API routes
	"log"                // For logging errors

	"github.com/gin-gonic/gin" // Web framework for Go
	"github.com/joho/godotenv" // For loading environment variables
)

// The main function is where our application starts running
func main() {
	// Load environment variables from .env file
	// This helps us keep sensitive information like database passwords secure
	err := godotenv.Load()

	// If we can't load the .env file, stop the application
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Create a new Gin router
	// This will handle all our web requests
	router := gin.Default()

	// Create a simple test route
	// When someone visits the homepage ("/"), we send a welcome message
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Go Gin Auth! Server Working!",
		})
	})

	// Connect to our database using the configuration
	config.ConnectDB()

	// Create the User table in our database if it doesn't exist
	if err := config.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("User migration failed: %v", err)
	}

	// Create the Post table in our database if it doesn't exist
	if err := config.DB.AutoMigrate(&models.Post{}); err != nil {
		log.Fatalf("Post migration failed: %v", err)
	}

	// Set up all our API routes (like login, register, etc.)
	routes.SetupRoutes(router)

	// Start the web server on port 8080
	// This makes our application available to receive requests
	router.Run(":8080")
}
