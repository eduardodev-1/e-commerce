package ports

import (
	"e-commerce/internal/core/domain"
	"e-commerce/internal/error"
)

type UserService interface {
	GetUserRoles(username string) ([]string, error)
	Authenticate(credentials *domain.RequestCredentials) (*domain.AuthenticatedUser, error)
	GetPaginatedList(requestParams *domain.RequestParams, page *domain.Page) (*domain.Page, *http_error.ErrorParams)
	Get(id int) (*domain.User, *http_error.ErrorParams)
}
type ProductService interface {
	GetPaginatedList(requestParams *domain.RequestParams, page *domain.Page) (*domain.Page, *http_error.ErrorParams)
	Get(id int) (*domain.Product, *http_error.ErrorParams)
}
