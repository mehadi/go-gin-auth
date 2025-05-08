package seeder

import (
	"go-gin-auth-api-starter-kit/config"
	"go-gin-auth-api-starter-kit/models"
	"go-gin-auth-api-starter-kit/utils"
	"log"
)

// SeedUsers creates initial users if they don't exist
func SeedUsers() error {
	// Check if any users exist
	var count int64
	if err := config.DB.Model(&models.User{}).Count(&count).Error; err != nil {
		return err
	}

	// If users exist, don't seed
	if count > 0 {
		log.Println("Users already exist, skipping seeding")
		return nil
	}

	// Sample users to seed
	users := []models.User{
		{
			Username: "admin",
			Email:    "admin@example.com",
			Password: utils.HashPasswordOrPanic("admin123"),
		},
		{
			Username: "user1",
			Email:    "user1@example.com",
			Password: utils.HashPasswordOrPanic("user123"),
		},
		{
			Username: "user2",
			Email:    "user2@example.com",
			Password: utils.HashPasswordOrPanic("user123"),
		},
	}

	// Create users
	for _, user := range users {
		if err := config.DB.Create(&user).Error; err != nil {
			return err
		}
		log.Printf("Seeded user: %s", user.Username)
	}

	log.Println("Successfully seeded users")
	return nil
}

// ForceSeedUsers seeds users regardless of whether they exist
func ForceSeedUsers() error {
	// Delete all existing users
	if err := config.DB.Where("1 = 1").Delete(&models.User{}).Error; err != nil {
		return err
	}

	// Seed new users
	return SeedUsers()
}
