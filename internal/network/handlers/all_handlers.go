package handlers

import (
	"e-commerce/internal/core/adapters/services"
)

type Handlers struct {
	UserHandler    *UserHandler
	LoginHandler   *LoginHandler
	ProductHandler *ProductHandler
}

func NewHandlers(allServices *services.Services) *Handlers {
	return &Handlers{
		UserHandler:    NewUserHandler(allServices.UserService),
		LoginHandler:   NewLoginHandler(allServices.UserService),
		ProductHandler: NewProductHandler(allServices.ProductService),
	}
}
