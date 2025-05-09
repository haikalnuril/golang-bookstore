// internal/app/utils/jwt.go
package utils

import (
	"bookstore/internal/app/config"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var cfg = config.LoadConfig()

// JWT secret key - in production, store this in an environment variable
var jwtSecretKey = []byte(cfg.JWTSecret)

// JWTClaims represents the claims in the JWT token
type JWTClaims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	jwt.RegisteredClaims
}

// GenerateJWTToken creates a new JWT token for a user
func GenerateJWTToken(userID, email, name string) (string, error) {
	// Set token expiration time (e.g., 24 hours)
	expirationTime := time.Now().Add(time.Duration(cfg.JWTExpire) * 24 * time.Hour)

	// Create claims with user data
	claims := &JWTClaims{
		UserID: userID,
		Email:  email,
		Name:   name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "bookstore-api",
			Subject:   userID,
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWTToken validates the JWT token and returns the claims
func ValidateJWTToken(tokenString string) (*JWTClaims, error) {
	// Parse and validate the token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JWTClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecretKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	// Extract claims
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
