package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Secret key - in production, use environment variables
var secretKey = []byte("your-secret-key")

// Custom claims - add whatever you need here
type CustomClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken creates a new JWT token
func GenerateToken(userID, email string) (string, error) {
	claims := CustomClaims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "your-app-name",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

// ValidateToken verifies the JWT token
func ValidateToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func main() {
	// Example usage
	token, err := GenerateToken("123", "user@example.com")
	if err != nil {
		fmt.Println("Error generating token:", err)
		return
	}

	fmt.Println("Generated Token:", token)

	// Validate the token
	claims, err := ValidateToken(token)
	if err != nil {
		fmt.Println("Error validating token:", err)
		return
	}

	fmt.Printf("Valid Token Claims: %+v\n", claims)
}
