package handlers

import (
	"e-commerce/internal/core/adapters/services"
)

type Handlers struct {
	UserHandler     *UserHandler
	LoginHandler    *AuthHandler
	ProductHandler  *ProductHandler
	CategoryHandler *CategoryHandler
}

func NewHandlers(allServices *services.Services) *Handlers {
	return &Handlers{
		UserHandler:     NewUserHandler(allServices.UserService),
		LoginHandler:    NewAuthHandler(allServices.UserService),
		ProductHandler:  NewProductHandler(allServices.ProductService),
		CategoryHandler: NewCategoryHandler(allServices.CategoryService),
	}
}
