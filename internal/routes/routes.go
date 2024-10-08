package routes

import (
	"e-commerce/internal/core/handlers"
	"e-commerce/internal/infra/middleware"
	"github.com/gofiber/fiber/v2"
)

func Public(app *fiber.App, allHandlers *handlers.Handlers) {
	v1 := app.Group("/v1")
	login := v1.Group("/login")
	{
		login.Post("/auth", allHandlers.LoginHandler.Authenticate)
		login.Post("/new-user", allHandlers.LoginHandler.Post)
	}
	product := v1.Group("/product")
	{
		product.Get("/", allHandlers.ProductHandler.GetPaginatedList)
		product.Get("/:id", allHandlers.ProductHandler.Get)
	}
	categories := v1.Group("/category")
	{
		categories.Get("/", allHandlers.CategoryHandler.GetPaginatedList)
		categories.Get("/:id", allHandlers.CategoryHandler.Get)
	}
}

func Private(app *fiber.App, allHandlers *handlers.Handlers) {
	app.Use(middleware.IsAuthenticatedMiddleware)
	v1 := app.Group("/v1")
	user := v1.Group("/user")
	{
		user.Get("/", middleware.AuthorizationMiddleware("ROLE_ADMIN"), allHandlers.UserHandler.GetPaginatedList)
		user.Get("/me", allHandlers.UserHandler.GetMe)
		user.Put("/me", allHandlers.UserHandler.UpdateMe)
		user.Delete("/me", allHandlers.UserHandler.DeleteMe)
		user.Get("/:id", middleware.AuthorizationMiddleware("ROLE_ADMIN"), allHandlers.UserHandler.Get)
		user.Put("/:id", middleware.AuthorizationMiddleware("ROLE_ADMIN"), allHandlers.UserHandler.Update)
		user.Delete("/:id", middleware.AuthorizationMiddleware("ROLE_ADMIN"), allHandlers.UserHandler.Delete)
	}
	product := v1.Group("/product")
	{
		product.Post("/", middleware.AuthorizationMiddleware("ROLE_ADMIN", "ROLE_SELLER"), allHandlers.ProductHandler.Post)
		product.Put("/:id", middleware.AuthorizationMiddleware("ROLE_ADMIN", "ROLE_SELLER"), allHandlers.ProductHandler.Update)
		product.Delete("/:id", middleware.AuthorizationMiddleware("ROLE_ADMIN", "ROLE_SELLER"), allHandlers.ProductHandler.Delete)
	}
	categories := v1.Group("/category")
	{
		categories.Post("/", middleware.AuthorizationMiddleware("ROLE_ADMIN", "ROLE_SELLER"), allHandlers.CategoryHandler.Post)
		categories.Put("/:id", middleware.AuthorizationMiddleware("ROLE_ADMIN"), allHandlers.CategoryHandler.Update)
		categories.Delete("/:id", middleware.AuthorizationMiddleware("ROLE_ADMIN"), allHandlers.CategoryHandler.Delete)
	}
	//orders
	//payment
}
