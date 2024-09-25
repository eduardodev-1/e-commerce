package ports

import (
	"e-commerce/internal/core/domain/models"
	httpError "e-commerce/internal/httperror"
)

type ProductRepository interface {
	FindPaginatedWithTotalCount(params *models.QueryParams) (*[]models.Product, int, *httpError.ErrorParams)
	FindById(id int) (*models.Product, *httpError.ErrorParams)
	Insert(product *models.Product, username string, isAdmin bool) (id int, errorParams *httpError.ErrorParams)
	Update(product *models.Product, username string, isAdmin bool) *httpError.ErrorParams
	Delete(id int, username string, isAdmin bool) *httpError.ErrorParams
}
type CategoryRepository interface {
	FindPaginatedWithTotalCount(params *models.QueryParams) (*[]models.Category, int, *httpError.ErrorParams)
	FindById(id int) (*models.Category, *httpError.ErrorParams)
	Insert(category *models.Category) (id int, errorParams *httpError.ErrorParams)
	Update(category *models.Category) *httpError.ErrorParams
	Delete(id int) *httpError.ErrorParams
}
type UserRepository interface {
	GetAuthoritiesByUsername(username string) ([]string, error)
	GetAuthenticationData(username string) (user *models.AuthenticatedUser, hashedpassword string, errorParams *httpError.ErrorParams)
	FindPaginatedWithTotalCount(params *models.QueryParams) (*[]models.User, int, *httpError.ErrorParams)
	FindById(id int) (*models.User, *httpError.ErrorParams)
	Insert(*models.UserFromRequest) (id int, errorParams *httpError.ErrorParams)
	FindByUsername(userName string) (*models.User, *httpError.ErrorParams)
	Update(update *models.UserUpdateRequest) *httpError.ErrorParams
	Delete(id int) *httpError.ErrorParams
}
