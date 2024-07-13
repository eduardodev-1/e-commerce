package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"strings"
)

func GetFiberConfig() *fiber.App {
	app := fiber.New()
	app.Use(
		cors.New(
			cors.Config{
				AllowOrigins: "*",
				AllowMethods: strings.Join([]string{
					fiber.MethodGet,
					fiber.MethodPost,
					fiber.MethodHead,
					fiber.MethodPut,
					fiber.MethodDelete,
					fiber.MethodPatch,
					fiber.MethodOptions,
				}, ","),
			},
		),
		recover.New(recover.Config{EnableStackTrace: true}),
		func(c *fiber.Ctx) error {
			if strings.Contains(c.OriginalURL(), "/favicon.ico") {
				return c.Next() // Skip logging for favicon requests
			}
			return logger.New()(c)
		},
	)
	return app
}
