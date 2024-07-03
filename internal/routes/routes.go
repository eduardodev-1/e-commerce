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
}

func PrivateRoutes(app *fiber.App, allControllers *controller.Controllers) {
	v1 := app.Group("/v1")
	_ = v1.Group("/usuario")
	{
		//u.Get("/", allControllers.UsuarioController.List)
		//u.Get("/create", allControllers.UsuarioController.Create)
		//u.Get("/:id", allControllers.UsuarioController.Get)
		//u.Post("/", allControllers.UsuarioController.Save)
		//u.Get("/edit/:id", allControllers.UsuarioController.Edit)
		//u.Put("/:id", allControllers.UsuarioController.Update)
		//u.Delete("/:id", allControllers.UsuarioController.Delete)
	}
}
