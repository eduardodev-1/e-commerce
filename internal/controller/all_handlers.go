package controller

import (
	"github.com/eduardodev-1/e-commerce/internal/services"
)

type Controllers struct {
	UsuarioController UsuarioController
	LoginController   LoginController
	//others Controllers
}

func NewControllers(allServices *services.Services) *Controllers {
	return &Controllers{
		UsuarioController: NewUsuarioController(allServices.UsuarioService),
		LoginController:   NewLoginController(allServices.UsuarioService),
		//others Controllers
	}
}
