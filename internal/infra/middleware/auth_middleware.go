package middleware

import (
	"e-commerce/internal/auth"
	"strings"

	"github.com/gofiber/fiber/v2"
	_ "github.com/golang-jwt/jwt"
)

func IsAuthenticatedMiddleware(ctx *fiber.Ctx) error {
	headerAuth := strings.SplitAfter(ctx.Get("Authorization"), "Bearer ")
	var tokenString = ""
	if len(headerAuth) == 2 {
		tokenString = headerAuth[1]
	}
	if tokenString == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "No token provided",
		})
	}
	token, err := auth.ParseToken(tokenString)
	if err != nil {
		return ctx.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	if !token.Valid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid token",
		})
	}
	username, authorities, userId, err := auth.ValidateAndExtractTokenData(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	err = auth.ValidateRouteAuthority(username, authorities, ctx.Route(), ctx.Path())
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Token is valid, add user information to context
	ctx.Locals("username", username)
	ctx.Locals("authorities", authorities)
	ctx.Locals("userId", userId)
	return ctx.Next()
}
