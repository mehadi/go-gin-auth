// This package handles database configuration and connection
package config

// Import necessary packages
import (
	"fmt" // For string formatting
	"log" // For logging errors
	"os"  // For environment variables

	"github.com/joho/godotenv" // For loading .env file
	"gorm.io/driver/postgres"  // PostgreSQL driver for GORM
	"gorm.io/gorm"             // GORM ORM library
)

// DB is a global variable that holds our database connection
var DB *gorm.DB

// ConnectDB establishes a connection to the PostgreSQL database
func ConnectDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database connection details from environment variables
	dbHost := os.Getenv("DB_HOST")         // Database server address
	dbPort := os.Getenv("DB_PORT")         // Database server port
	dbUser := os.Getenv("DB_USER")         // Database username
	dbPassword := os.Getenv("DB_PASSWORD") // Database password
	dbName := os.Getenv("DB_NAME")         // Database name

	// Create the database connection string
	// This string contains all the information needed to connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// Open a connection to the database using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// If connection fails, stop the application
		log.Fatal("Failed to connect to DB: ", err)
	}

	// Store the database connection in our global variable
	DB = db
	fmt.Println("Database connection established")
}
