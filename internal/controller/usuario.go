package controller

import (
	"e-commerce/internal/services"
)

type UsuarioController struct {
	UsuarioService services.UsuarioService
}

func NewUsuarioController(usuarioSvc services.UsuarioService) UsuarioController {
	return UsuarioController{
		UsuarioService: usuarioSvc,
	}
}
