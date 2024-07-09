package controller

import (
	"e-commerce/internal/services"
)

type Controllers struct {
	UserController     *UserController
	LoginController    *LoginController
	ProductController  *ProductController
	DataBaseController *DataBaseController
}

func NewControllers(allServices *services.Services) *Controllers {
	return &Controllers{
		UserController:     NewUserController(allServices.UserService),
		LoginController:    NewLoginController(allServices.UserService),
		ProductController:  NewProductController(allServices.ProductService),
		DataBaseController: NewDataBaseController(allServices.DataBaseService),
	}
}
