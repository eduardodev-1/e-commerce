package ports

import (
	"e-commerce/internal/core/domain"
	httpError "e-commerce/internal/error"
)

type ProductRepository interface {
	FindPaginatedWithTotalCount(params *domain.QueryParams) (*[]domain.Product, domain.TotalCount, *httpError.ErrorParams)
	FindById(id int) (*domain.Product, *httpError.ErrorParams)
	Insert(product *domain.Product) (domain.IdToResponse, *httpError.ErrorParams)
}
type UserRepository interface {
	GetAuthoritiesByUserName(username string) ([]string, error)
	GetAuthenticationData(username string) (*domain.AuthenticatedUser, string, error)
	FindPaginatedWithTotalCount(params *domain.QueryParams) (*[]domain.User, domain.TotalCount, *httpError.ErrorParams)
	FindById(id int) (*domain.User, *httpError.ErrorParams)
	Insert(user *domain.User) (*domain.User, *httpError.ErrorParams)
}
