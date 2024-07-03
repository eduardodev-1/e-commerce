package middleware

import (
	"e-commerce/internal/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-jwt/jwt"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Extract the JWT token from the Authorization header
	headerAuth := strings.SplitAfter(c.Get("Authorization"), "Bearer ")
	var tokenString = ""
	if len(headerAuth) == 2 {
		tokenString = headerAuth[1]
	}
	if tokenString == "" {
		// No token provided, return unauthorized status
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}
	// Parse the JWT token
	token, err := auth.ParseToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Validate Token
	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}
	// Extract data from token
	username, authorities, err := auth.ValidateAndExtractTokenData(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Token is valid, add user information to context
	c.Locals("username", username)
	c.Locals("authorities", authorities)
	return c.Next()
}
