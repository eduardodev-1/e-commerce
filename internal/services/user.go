package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"e-commerce/internal/utils"
)

type UserService struct {
	UsuarioRepo *repositories.UserRepository
}

func NewUserService(usuarioRepo *repositories.UserRepository) *UserService {
	return &UserService{
		UsuarioRepo: usuarioRepo,
	}
}
func (s *UserService) GetUserRoles(username string) ([]string, error) {
	authorities, err := s.UsuarioRepo.GetAuthoritiesByUserName(username)
	return authorities, err

}
func (s *UserService) Authenticate(credentials *models.RequestCredentials) (*models.AuthenticatedUser, error) {

	user, hashedPassword, err := s.UsuarioRepo.GetAuthenticationData(credentials.Username)
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
