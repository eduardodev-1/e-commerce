package tests

import "github.com/gofiber/fiber/v2"

func IsAuthenticatedMiddlewareMock(ctx *fiber.Ctx) error {
	// Token is valid, add user information to context
	ctx.Locals("username", "test@gmail.com")
	ctx.Locals("authorities", []string{"ROLE_ADMIN"})
	ctx.Locals("userId", "4")
	return ctx.Next()
}
