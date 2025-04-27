package main

import (
	"flag"
	"go-gin-auth/config"
	"go-gin-auth/pkg/seeder"
	"log"
)

func main() {
	// Parse command line flags
	force := flag.Bool("force", false, "Force reseed by deleting existing users")
	flag.Parse()

	// Connect to database
	config.ConnectDB()

	// Run seeder
	var err error
	if *force {
		err = seeder.ForceSeedUsers()
	} else {
		err = seeder.SeedUsers()
	}

	if err != nil {
		log.Fatalf("Error seeding users: %v", err)
	}
}
