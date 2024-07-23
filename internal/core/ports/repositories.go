package ports

import (
	"e-commerce/internal/core/domain/models"
	httpError "e-commerce/internal/error"
)

type ProductRepository interface {
	FindPaginatedWithTotalCount(params *models.QueryParams) (*[]models.Product, models.TotalCount, *httpError.ErrorParams)
	FindById(id int) (*models.Product, *httpError.ErrorParams)
	Insert(product *models.Product) (id int, errorParams *httpError.ErrorParams)
}
type UserRepository interface {
	GetAuthoritiesByUserName(username string) ([]string, error)
	GetAuthenticationData(username string) (user *models.AuthenticatedUser, hashedpassword string, err error)
	FindPaginatedWithTotalCount(params *models.QueryParams) (*[]models.User, models.TotalCount, *httpError.ErrorParams)
	FindById(id int) (*models.User, *httpError.ErrorParams)
	Insert(*models.UserFromRequest) (id int, errorParams *httpError.ErrorParams)
	FindByUserName(userName string) (*models.User, *httpError.ErrorParams)
	Update(update *models.UserFromRequest) *httpError.ErrorParams
}
