package routes

import (
	"e-commerce/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Public(app *fiber.App, allHandlers *handler.Handlers) {
	v1 := app.Group("/v1")
	login := v1.Group("/login")
	{
		login.Post("/submit", allHandlers.LoginHandler.Authenticate)
	}
	product := v1.Group("/product")
	{
		product.Get("/", allHandlers.ProductHandler.GetPaginatedList)
		product.Get("/:id", allHandlers.ProductHandler.Get)
	}
}

func Private(app *fiber.App, allHandlers *handler.Handlers) {
	v1 := app.Group("/v1")
	user := v1.Group("/user")
	{
		user.Get("/", allHandlers.UserHandler.GetPaginatedList)
		user.Get("/:id", allHandlers.UserHandler.Get)
	}
	product := v1.Group("/product")
	{
		product.Get("/", allHandlers.ProductHandler.GetPaginatedList)
		product.Get("/:id", allHandlers.ProductHandler.Get)
	}
}
