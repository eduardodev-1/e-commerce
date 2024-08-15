package handlers

import (
	"e-commerce/internal/core/adapters/handlers/fiber"
	"e-commerce/internal/core/domain/services"
)

type Handlers struct {
	UserHandler     *fiber.UserHandler
	LoginHandler    *fiber.AuthHandler
	ProductHandler  *fiber.ProductHandler
	CategoryHandler *fiber.CategoryHandler
}

func NewHandlers(allServices *services.Services) *Handlers {
	return &Handlers{
		UserHandler:     fiber.NewUserHandler(allServices.UserService),
		LoginHandler:    fiber.NewAuthHandler(allServices.UserService),
		ProductHandler:  fiber.NewProductHandler(allServices.ProductService),
		CategoryHandler: fiber.NewCategoryHandler(allServices.CategoryService),
	}
}
