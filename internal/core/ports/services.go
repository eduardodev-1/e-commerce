package ports

import (
	"e-commerce/internal/core/domain"
	httpError "e-commerce/internal/error"
)

type UserService interface {
	GetUserRoles(username string) ([]string, error)
	Authenticate(credentials *domain.RequestCredentials) (*domain.AuthenticatedUser, error)
	GetPaginatedList(requestParams *domain.RequestParams) (*domain.Page, *httpError.ErrorParams)
	Get(id int) (*domain.User, *httpError.ErrorParams)
}
type ProductService interface {
	GetPaginatedList(requestParams *domain.RequestParams) (*domain.Page, *httpError.ErrorParams)
	Get(id int) (*domain.Product, *httpError.ErrorParams)
	Post(*domain.Product) (domain.IdToResponse, *httpError.ErrorParams)
}
