package services

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/core/domain"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/error"
	"github.com/gofiber/fiber/v2"
	"strconv"
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
func (s *UserService) AuthenticateUserWithPasswordCredentials(credentials *domain.RequestCredentials) (*domain.LoginResponse, error) {

	user, hashedPassword, err := s.UserRepository.GetAuthenticationData(credentials.Username)
	if err != nil {
		return nil, err
	}
	passwordPair := domain.PasswordPair{
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
	// Create JWToken
	token, err := auth.NewJWToken(user.Id, user.Username, roles)
	if err != nil {
		return nil, err
	}
	return &domain.LoginResponse{Token: token}, nil
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
func (s *UserService) Get(id string, userName string) (*domain.User, *httpError.ErrorParams) {
	if id == "me" {
		user, errorParams := s.UserRepository.FindByUserName(userName)
		if errorParams != nil {
			return nil, errorParams
		}
		return user, nil
	}
	idToInt, err := strconv.Atoi(id)
	if err != nil {
		return nil, &httpError.ErrorParams{
			Message: "invalid id",
			Status:  fiber.StatusBadRequest,
		}
	}
	user, errorParams := s.UserRepository.FindById(idToInt)
	if errorParams != nil {
		return nil, errorParams
	}
	return user, nil
}
func (s *UserService) CreateNewUser(newUser *domain.NewUserRequest) (int, *httpError.ErrorParams) {
	errorParams := new(httpError.ErrorParams)
	errorParams = newUser.CheckUserType()
	if errorParams != nil {
		return 0, errorParams
	}
	errorParams = newUser.SetEncryptedPassword()
	if errorParams != nil {
		return 0, errorParams
	}
	newUserId, errorParams := s.UserRepository.Insert(newUser)
	return newUserId, errorParams
}
