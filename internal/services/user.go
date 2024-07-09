package services

import (
	fiber_error "e-commerce/internal/error"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"e-commerce/internal/utils"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(usuarioRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: usuarioRepo,
	}
}
func (s *UserService) GetUserRoles(username string) ([]string, error) {
	authorities, err := s.UserRepository.GetAuthoritiesByUserName(username)
	return authorities, err

}
func (s *UserService) Authenticate(credentials *models.RequestCredentials) (*models.AuthenticatedUser, error) {

	user, hashedPassword, err := s.UserRepository.GetAuthenticationData(credentials.Username)
	if err != nil {
		return nil, err
	}
	passwordPair := models.PasswordPair{
		Password:       credentials.Password,
		HashedPassword: hashedPassword,
	}
	if err = utils.CheckPasswordRequest(passwordPair); err != nil {
		return nil, err
	}

	return user, nil
}
func (s UserService) GetPaginatedList(requestParams *models.RequestParams, page *models.Page) (*models.Page, *fiber_error.ErrorParams) {
	page.SetRequestParams(requestParams)
	queryParams := page.GetQueryParams()
	content, count, errorParams := s.UserRepository.PageableFindAll(queryParams)
	if errorParams != nil {
		return nil, errorParams
	}
	page.SetResultParams(content, count)
	return page, nil
}

func (s UserService) Get(id int) (*models.User, *fiber_error.ErrorParams) {
	product, errorParams := s.UserRepository.GetById(id)
	if errorParams != nil {
		return nil, errorParams
	}
	return product, nil
}
