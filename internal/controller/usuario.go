package controller

import (
	"github.com/eduardodev-1/e-commerce/internal/services"
)

type UsuarioController struct {
	UsuarioService services.UsuarioService
}

func NewUsuarioController(usuarioSvc services.UsuarioService) UsuarioController {
	return UsuarioController{
		UsuarioService: usuarioSvc,
	}
}
