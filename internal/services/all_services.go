package services

import (
	"e-commerce/internal/repositories"
)

type Services struct {
	UserService     *UserService
	ProductService  *ProductService
	DataBaseService *DataBaseService
}

func NewServices(allRepositories *repositories.Repositories) *Services {
	return &Services{
		UserService:     NewUserService(allRepositories.UserRepository),
		ProductService:  NewProductService(allRepositories.ProductRepository),
		DataBaseService: NewDataBaseService(allRepositories.DataBaseRepository),
	}
}
