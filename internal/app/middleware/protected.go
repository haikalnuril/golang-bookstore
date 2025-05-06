package middleware

import (
	"bookstore/internal/app/exception"
	"bookstore/internal/app/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get authorization header
		authHeader := c.Get("Authorization")

		// Check if the token exists in the header
		if authHeader == "" {
			return &exception.UnauthorizedError{Message: "Missing authorization token"}
		}

		// Check if the format is correct (Bearer token)
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return &exception.UnauthorizedError{Message: "Invalid token format"}
		}

		// Extract the token
		tokenString := parts[1]

		// Validate the token
		claims, err := utils.ValidateJWTToken(tokenString)
		if err != nil {
			return &exception.UnauthorizedError{Message: "Invalid or expired token"}
		}

		// Set user details in locals for use in handlers
		c.Locals("userID", claims.UserID)
		c.Locals("email", claims.Email)
		c.Locals("name", claims.Name)

		// Continue to the next middleware/handler
		return c.Next()
	}
}
