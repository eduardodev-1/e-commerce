package routes

import (
	"e-commerce/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App, allControllers *controller.Controllers) {
	v1 := app.Group("/v1")
	dataBase := v1.Group("/database")
	{
		dataBase.Put("/restart", allControllers.DataBaseController.ResetDataBase)
	}
	login := v1.Group("/login")
	{
		login.Post("/submit", allControllers.LoginController.Authenticate)
	}
	product := v1.Group("/product")
	{
		product.Get("/", allControllers.ProductController.GetPaginatedList)
		product.Get("/:id", allControllers.ProductController.Get)
	}
}

func PrivateRoutes(app *fiber.App, allControllers *controller.Controllers) {
	v1 := app.Group("/v1")
	user := v1.Group("/user")
	{
		user.Get("/", allControllers.UserController.GetPaginatedList)
		user.Get("/:id", allControllers.UserController.Get)
	}
	product := v1.Group("/product")
	{
		product.Get("/", allControllers.ProductController.GetPaginatedList)
		product.Get("/:id", allControllers.ProductController.Get)
	}
}
