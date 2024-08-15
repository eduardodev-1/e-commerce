package services

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/httperror"
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
	authorities, err := s.UserRepository.GetAuthoritiesByUsername(username)
	return authorities, err

}
func (s *UserService) AuthenticateUserWithPasswordCredentials(credentials *models.RequestCredentials) (*models.LoginResponse, *httpError.ErrorParams) {
	user, hashedPassword, errorParams := s.UserRepository.GetAuthenticationData(credentials.Username)
	if errorParams != nil {
		return nil, errorParams
	}
	passwordPair := models.PasswordPair{
		OriginalPassword: credentials.Password,
		HashedPassword:   hashedPassword,
	}
	if err := passwordPair.CheckRequestPassword(); err != nil {
		errorParams = &httpError.ErrorParams{}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	roles, err := s.GetUserRoles(user.Username)
	if err != nil {
		errorParams = &httpError.ErrorParams{}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	token, err := auth.NewJWToken(user.Id, user.Username, roles)
	if err != nil {
		errorParams = &httpError.ErrorParams{}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
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
	errorParams = utils.CheckUserType(newUser.UserType)
	if errorParams != nil {
		return 0, errorParams
	}
	newUser.User.Login.Password, errorParams = utils.EncryptPassword(newUser.User.Login.Password)
	if errorParams != nil {
		return 0, errorParams
	}
	newUserId, errorParams := s.UserRepository.Insert(newUser)
	return newUserId, errorParams
}
func (s *UserService) Update(username string, userToUpdate *models.UserUpdateRequest) *httpError.ErrorParams {
	errorParams := new(httpError.ErrorParams)
	errorParams = utils.CheckUserType(userToUpdate.UserType)
	if errorParams != nil {
		return errorParams
	}
	errorParams = utils.CheckUsername(userToUpdate.User.Login.Username, username)
	if errorParams != nil {
		return errorParams
	}
	_, hashedPassword, errorParams := s.UserRepository.GetAuthenticationData(username)
	if errorParams != nil {
		return errorParams
	}
	passwordPair := models.PasswordPair{
		OriginalPassword: userToUpdate.User.Login.Password,
		HashedPassword:   hashedPassword,
	}
	if err := passwordPair.CheckRequestPassword(); err != nil {
		errorParams = &httpError.ErrorParams{}
		errorParams.SetDefaultParams(err)
		return errorParams
	}
	if userToUpdate.PasswordFields.UpdatePassword {
		userToUpdate.User.Login.Password = userToUpdate.PasswordFields.NewPassword
	}
	if userToUpdate.UserFields.UpdateUsername {
		userToUpdate.User.Login.Username = userToUpdate.UserFields.NewUsername
	}
	userToUpdate.User.Login.Password, errorParams = utils.EncryptPassword(userToUpdate.User.Login.Password)
	if errorParams != nil {
		return errorParams
	}
	errorParams = s.UserRepository.Update(userToUpdate)
	if errorParams != nil {
		return errorParams
	}
	return nil
}
func (s *UserService) Delete(id int) *httpError.ErrorParams {
	errorParams := s.UserRepository.Delete(id)
	if errorParams != nil {
		return errorParams
	}
	return nil
}
