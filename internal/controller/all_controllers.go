package controller

import (
	"e-commerce/internal/services"
)

type Controllers struct {
	UsuarioController  *UserController
	LoginController    *LoginController
	ProductController  *ProductController
	DataBaseController *DataBaseController
}

func NewControllers(allServices *services.Services) *Controllers {
	return &Controllers{
		UsuarioController:  NewUserController(allServices.UserService),
		LoginController:    NewLoginController(allServices.UserService),
		ProductController:  NewProductController(allServices.ProductService),
		DataBaseController: NewDataBaseController(allServices.DataBaseService),
	}
}
