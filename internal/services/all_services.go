package services

import (
	"github.com/eduardodev-1/e-commerce/internal/repositories"
)

type Services struct {
	UsuarioService UsuarioService
}

func NewServices(allRepositories *repositories.Repositories) *Services {
	return &Services{
		UsuarioService: NewUsuarioService(
			allRepositories.UsuarioRepository,
		),
	}
}
