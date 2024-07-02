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
				AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
			},
		),
		//Configurar recover
		recover.New(recover.Config{EnableStackTrace: true}),
		// Exclude favicon requests from logging
		func(c *fiber.Ctx) error {
			if strings.Contains(c.OriginalURL(), "/favicon.ico") {
				return c.Next() // Skip logging for favicon requests
			}
			return logger.New()(c) // Log all other requests
		},
	)
	return app
}
