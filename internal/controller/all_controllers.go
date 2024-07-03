package controller

import (
	"e-commerce/internal/services"
)

type Controllers struct {
	UsuarioController UsuarioController
	LoginController   LoginController
	ProductController ProductController
}

func NewControllers(allServices *services.Services) *Controllers {
	return &Controllers{
		UsuarioController: NewUsuarioController(allServices.UserService),
		LoginController:   NewLoginController(allServices.UserService),
		ProductController: NewProductController(allServices.ProductService),
	}
}
