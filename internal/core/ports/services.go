package ports

import (
	"e-commerce/internal/core/domain"
	httpError "e-commerce/internal/error"
)

type UserService interface {
	GetUserRoles(username string) ([]string, error)
	AuthenticateUserWithPasswordCredentials(credentials *domain.RequestCredentials) (*domain.LoginResponse, error)
	GetPaginatedList(requestParams *domain.RequestParams) (*domain.Page, *httpError.ErrorParams)
	Get(id string, userName string) (*domain.User, *httpError.ErrorParams)
	CreateNewUser(newUser *domain.NewUserRequest) (id int, errorParams *httpError.ErrorParams)
}
type ProductService interface {
	GetPaginatedList(requestParams *domain.RequestParams) (*domain.Page, *httpError.ErrorParams)
	Get(id int) (*domain.Product, *httpError.ErrorParams)
	Post(*domain.Product) (id int, errorParams *httpError.ErrorParams)
}
