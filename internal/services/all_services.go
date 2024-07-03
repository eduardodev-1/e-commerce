package services

import (
	"e-commerce/internal/repositories"
)

type Services struct {
	UserService    UsuarioService
	ProductService ProductService
}

func NewServices(allRepositories *repositories.Repositories) *Services {
	return &Services{
		UserService:    NewUserService(allRepositories.UserRepository),
		ProductService: NewProductService(allRepositories.ProductRepository),
	}
}
