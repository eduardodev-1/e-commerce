package ports

import (
	"e-commerce/internal/core/domain/models"
	httpError "e-commerce/internal/httperror"
)

type UserService interface {
	GetUserRoles(username string) ([]string, error)
	AuthenticateUserWithPasswordCredentials(credentials *models.RequestCredentials) (*models.LoginResponse, *httpError.ErrorParams)
	GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams)
	Get(id int) (*models.User, *httpError.ErrorParams)
	CreateNewUser(newUser *models.UserFromRequest) (id int, errorParams *httpError.ErrorParams)
	Update(username string, userToUpdate *models.UserUpdateRequest) *httpError.ErrorParams
	Delete(id int) *httpError.ErrorParams
}
type ProductService interface {
	GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams)
	Get(id int) (*models.Product, *httpError.ErrorParams)
	Post(product *models.Product, username string, isAdmin bool) (id int, errorParams *httpError.ErrorParams)
	Update(product *models.Product, username string, isAdmin bool) *httpError.ErrorParams
	Delete(id int, username string, isAdmin bool) *httpError.ErrorParams
}
type CategoryService interface {
	GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams)
	Get(id int) (*models.Category, *httpError.ErrorParams)
	Post(category *models.Category) (id int, errorParams *httpError.ErrorParams)
	Update(category *models.Category) *httpError.ErrorParams
	Delete(id int) *httpError.ErrorParams
}
