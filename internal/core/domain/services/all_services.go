package services

import (
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/ports"
)

type Services struct {
	UserService     ports.UserService
	ProductService  ports.ProductService
	CategoryService ports.CategoryService
}

func NewServices(allRepositories *repositories.Repositories) *Services {
	return &Services{
		UserService:     NewUserService(allRepositories.UserRepository),
		ProductService:  NewProductService(allRepositories.ProductRepository),
		CategoryService: NewCategoryService(allRepositories.CategoryRepository),
	}
}
