package services

import (
	"e-commerce/internal/core/domain"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/error"
	"e-commerce/internal/utils"
)

type UserService struct {
	UserRepository ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepo,
	}
}
func (s *UserService) GetUserRoles(username string) ([]string, error) {
	authorities, err := s.UserRepository.GetAuthoritiesByUserName(username)
	return authorities, err

}
func (s *UserService) Authenticate(credentials *domain.RequestCredentials) (*domain.AuthenticatedUser, error) {

	user, hashedPassword, err := s.UserRepository.GetAuthenticationData(credentials.Username)
	if err != nil {
		return nil, err
	}
	passwordPair := domain.PasswordPair{
		Password:       credentials.Password,
		HashedPassword: hashedPassword,
	}
	if err = utils.CheckPasswordRequest(passwordPair); err != nil {
		return nil, err
	}

	return user, nil
}
func (s *UserService) GetPaginatedList(requestParams *domain.RequestParams) (*domain.Page, *httpError.ErrorParams) {
	page := new(domain.Page)
	page.SetRequestParams(requestParams)
	queryParams := page.GetQueryParams()
	content, count, errorParams := s.UserRepository.FindPaginatedWithTotalCount(queryParams)
	if errorParams != nil {
		return nil, errorParams
	}
	page.SetResultParams(content, count)
	return page, nil
}

func (s *UserService) Get(id int) (*domain.User, *httpError.ErrorParams) {
	product, errorParams := s.UserRepository.FindById(id)
	if errorParams != nil {
		return nil, errorParams
	}
	return product, nil
}
