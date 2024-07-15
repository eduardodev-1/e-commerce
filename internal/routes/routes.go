package routes

import (
	"e-commerce/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Public(app *fiber.App, allHandlers *handler.Handlers) {
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
}

func Private(app *fiber.App, allHandlers *handler.Handlers) {
	v1 := app.Group("/v1")
	user := v1.Group("/user")
	{
		//ALL LOGGED
		user.Get("/me", allHandlers.UserHandler.GetMe)
		user.Put("/me", allHandlers.UserHandler.UpdateMe)
		user.Delete("/me", allHandlers.UserHandler.DeleteMe)
		//ROLE ADMIN ONLY
		user.Get("/", allHandlers.UserHandler.GetPaginatedList)
		user.Get("/:id", allHandlers.UserHandler.Get)
		user.Post("/", allHandlers.UserHandler.Post)
		user.Put("/:id", allHandlers.UserHandler.Update)
		user.Delete("/:id", allHandlers.UserHandler.Delete)
	}
	// ROLE_SELLER || ROLE_ADMIN
	//if seller, just him products.
	product := v1.Group("/product")
	{
		product.Post("/", allHandlers.ProductHandler.Post)
		product.Put("/:id", allHandlers.ProductHandler.Update)
		product.Delete("/:id", allHandlers.ProductHandler.Delete)
	}
	//categories := v1.Group("/category")
	//{
	//	categories.Get("/", allHandlers.CategoryHandler.Get)
	//}
}
