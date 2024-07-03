package routes

import (
	"e-commerce/internal/controller"
	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(app *fiber.App, allControllers *controller.Controllers) {
	v1 := app.Group("/v1")
	login := v1.Group("/login")
	{
		login.Post("/submit", allControllers.LoginController.Autenticate)
	}
	product := v1.Group("/product")
	{
		product.Get("/", allControllers.ProductController.GetPaginatedList)
	}
	dataBase := v1.Group("/database")
	{
		dataBase.Put("/restart", allControllers.DataBaseController.ResetDataBase)
	}
}

func PrivateRoutes(app *fiber.App, allControllers *controller.Controllers) {
	v1 := app.Group("/v1")
	_ = v1.Group("/usuario")
	{
		//u.Get("/", allControllers.UserController.List)
		//u.Get("/create", allControllers.UserController.Create)
		//u.Get("/:id", allControllers.UserController.Get)
		//u.Post("/", allControllers.UserController.Save)
		//u.Get("/edit/:id", allControllers.UserController.Edit)
		//u.Put("/:id", allControllers.UserController.Update)
		//u.Delete("/:id", allControllers.UserController.Delete)
	}
}
