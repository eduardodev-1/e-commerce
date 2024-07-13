package ports

import (
	"e-commerce/internal/core/domain"
	"e-commerce/internal/error"
)

type ProductRepository interface {
	FindPaginatedWithTotalCount(params *domain.QueryParams) (*[]domain.Product, domain.TotalCount, *http_error.ErrorParams)
	GetById(id int) (*domain.Product, *http_error.ErrorParams)
	Insert(product *domain.Product) (*domain.Product, *http_error.ErrorParams)
}
type UserRepository interface {
	GetAuthoritiesByUserName(username string) ([]string, error)
	GetAuthenticationData(username string) (*domain.AuthenticatedUser, string, error)
	FindPaginatedWithTotalCount(params *domain.QueryParams) (*[]domain.User, domain.TotalCount, *http_error.ErrorParams)
	GetById(id int) (*domain.User, *http_error.ErrorParams)
	Insert(user *domain.User) (*domain.User, *http_error.ErrorParams)
}
