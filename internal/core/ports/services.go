package ports

import (
	"e-commerce/internal/core/domain/models"
	httpError "e-commerce/internal/error"
)

type UserService interface {
	GetUserRoles(username string) ([]string, error)
	AuthenticateUserWithPasswordCredentials(credentials *models.RequestCredentials) (*models.LoginResponse, error)
	GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams)
	Get(id int) (*models.User, *httpError.ErrorParams)
	CreateNewUser(newUser *models.UserFromRequest) (id int, errorParams *httpError.ErrorParams)
	Update(request *models.UserFromRequest) *httpError.ErrorParams
	Delete(id int) *httpError.ErrorParams
}
type ProductService interface {
	GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams)
	Get(id int) (*models.Product, *httpError.ErrorParams)
	Post(*models.Product) (id int, errorParams *httpError.ErrorParams)
}
