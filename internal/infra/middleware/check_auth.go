package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AuthorizationMiddleware(roles ...string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorities := c.Locals("authorities").([]string)
		for _, role := range roles {
			for _, authority := range authorities {
				if role == authority {
					return c.Next()
				}
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "Access denied",
		})
	}
}

func HasRole(userAuthorities []string, role string) bool {
	for _, authority := range userAuthorities {
		if role == authority {
			return true
		}
	}
	return false
}
