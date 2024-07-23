package services

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/error"
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
func (s *UserService) AuthenticateUserWithPasswordCredentials(credentials *models.RequestCredentials) (*models.LoginResponse, error) {
	user, hashedPassword, err := s.UserRepository.GetAuthenticationData(credentials.Username)
	if err != nil {
		return nil, err
	}
	passwordPair := models.PasswordPair{
		Password:       credentials.Password,
		HashedPassword: hashedPassword,
	}
	if err = passwordPair.CheckPasswordRequest(); err != nil {
		return nil, err
	}
	roles, err := s.GetUserRoles(user.Username)
	if err != nil {
		return nil, err
	}
	token, err := auth.NewJWToken(user.Id, user.Username, roles)
	if err != nil {
		return nil, err
	}
	return &models.LoginResponse{Token: token}, nil
}
func (s *UserService) GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams) {
	page := new(models.Page)
	page.SetRequestParams(requestParams)
	queryParams := page.GetQueryParams()
	content, count, errorParams := s.UserRepository.FindPaginatedWithTotalCount(queryParams)
	if errorParams != nil {
		return nil, errorParams
	}
	page.SetResultParams(content, count)
	return page, nil
}
func (s *UserService) Get(id int) (*models.User, *httpError.ErrorParams) {
	user, errorParams := s.UserRepository.FindById(id)
	if errorParams != nil {
		return nil, errorParams
	}
	return user, nil
}
func (s *UserService) CreateNewUser(newUser *models.UserFromRequest) (int, *httpError.ErrorParams) {
	errorParams := new(httpError.ErrorParams)
	errorParams = newUser.CheckUserType()
	if errorParams != nil {
		return 0, errorParams
	}
	errorParams = newUser.EncryptPassword()
	if errorParams != nil {
		return 0, errorParams
	}
	newUserId, errorParams := s.UserRepository.Insert(newUser)
	return newUserId, errorParams
}
func (s *UserService) Update(userToUpdate *models.UserFromRequest) *httpError.ErrorParams {
	errorParams := new(httpError.ErrorParams)
	errorParams = userToUpdate.CheckUserType()
	if errorParams != nil {
		return errorParams
	}
	errorParams = userToUpdate.EncryptPassword()
	if errorParams != nil {
		return errorParams
	}
	errorParams := s.UserRepository.Update(userToUpdate)
	return nil
}
func (s *UserService) Delete(id int) *httpError.ErrorParams {
	//TODO implement me
	panic("implement me")
}
