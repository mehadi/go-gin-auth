// This package contains utility functions for JWT token handling
package utils

// Import necessary packages
import (
	"os"   // For environment variables
	"time" // For token expiration

	"github.com/dgrijalva/jwt-go" // For JWT operations
)

// Load the JWT secret key from environment variables
// This key is used to sign and verify tokens
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

// Claims represents the data stored in the JWT token
type Claims struct {
	Username           string // The username of the authenticated user
	jwt.StandardClaims        // Standard JWT claims like expiration time
}

// GenerateJWT creates a new JWT token for an authenticated user
// username: The username to include in the token
// Returns: The signed JWT token and any error that occurred
func GenerateJWT(username string) (string, error) {
	// Set token to expire in 24 hours
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create the JWT claims, which includes the username and expiration time
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // Convert to Unix timestamp
		},
	}

	// Create the token with our claims and sign it with our secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ValidateToken parses and validates a JWT token string.
// Returns the claims if the token is valid, or an error otherwise.
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
