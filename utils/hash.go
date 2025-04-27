// This package contains utility functions for password hashing
package utils

// Import the bcrypt package for secure password hashing
import "golang.org/x/crypto/bcrypt"

// HashPassword converts a plain text password into a secure hash
// password: The plain text password to hash
// Returns: The hashed password and any error that occurred
func HashPassword(password string) (string, error) {
	// Generate a secure hash from the password
	// 14 is the cost factor - higher is more secure but slower
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a plain text password with a stored hash
// password: The plain text password to check
// hash: The stored password hash to compare against
// Returns: true if the password matches the hash, false otherwise
func CheckPasswordHash(password, hash string) bool {
	// Compare the provided password with the stored hash
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	// If there's no error, the passwords match
	return err == nil
}

// HashPasswordOrPanic hashes a password and panics if there's an error
// This is useful for seeding data where we want to fail fast if hashing fails
func HashPasswordOrPanic(password string) string {
	hash, err := HashPassword(password)
	if err != nil {
		panic(err)
	}
	return hash
}
