package services

import (
	"e-commerce/internal/core/ports"
	"e-commerce/internal/repositories"
)

type Services struct {
	UserService    ports.UserService
	ProductService ports.ProductService
}

func NewServices(allRepositories *repositories.Repositories) *Services {
	return &Services{
		UserService:    NewUserService(allRepositories.UserRepository),
		ProductService: NewProductService(allRepositories.ProductRepository),
	}
}
