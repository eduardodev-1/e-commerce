package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"e-commerce/internal/utils"
)

type UsuarioService struct {
	UsuarioRepo repositories.UserRepository
}

func NewUserService(usuarioRepo repositories.UserRepository) UsuarioService {
	return UsuarioService{
		UsuarioRepo: usuarioRepo,
	}
}
func (s *UsuarioService) GetUserRoles(username string) ([]string, error) {
	authorities, err := s.UsuarioRepo.GetAuthoritiesByUserName(username)
	return authorities, err

}
func (s *UsuarioService) Authenticate(credentials *models.RequestCredentials) (*models.AuthenticatedUser, error) {

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
