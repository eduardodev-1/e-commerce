package services

import (
	"github.com/eduardodev-1/e-commerce/internal/models"
	"github.com/eduardodev-1/e-commerce/internal/repositories"
	"github.com/eduardodev-1/e-commerce/internal/utils"
)

type UsuarioService struct {
	UsuarioRepo repositories.UsuarioRepository
}

func NewUsuarioService(usuarioRepo repositories.UsuarioRepository) UsuarioService {
	return UsuarioService{
		UsuarioRepo: usuarioRepo,
	}
}
func (s *UsuarioService) GetUserRoles(username string) ([]string, error) {
	authorities, err := s.UsuarioRepo.GetAuthoritiesByUserName(username)
	return authorities, err

}
func (s *UsuarioService) Authenticate(credentials *models.LoginRequest) (*models.AuthenticatedUser, error) {

	usuario, hashedPassword, err := s.UsuarioRepo.GetAuthenticationData(credentials.Username)
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

	return usuario, nil
}
