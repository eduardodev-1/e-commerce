package handlers

import (
	"e-commerce/internal/core/domain/services"
	fiber2 "e-commerce/internal/core/handlers/fiber"
)

type Handlers struct {
	UserHandler     *fiber2.UserHandler
	LoginHandler    *fiber2.AuthHandler
	ProductHandler  *fiber2.ProductHandler
	CategoryHandler *fiber2.CategoryHandler
}

func NewHandlers(allServices *services.Services) *Handlers {
	return &Handlers{
		UserHandler:     fiber2.NewUserHandler(allServices.UserService),
		LoginHandler:    fiber2.NewAuthHandler(allServices.UserService),
		ProductHandler:  fiber2.NewProductHandler(allServices.ProductService),
		CategoryHandler: fiber2.NewCategoryHandler(allServices.CategoryService),
	}
}
