package controller

import (
	"e-commerce/internal/services"
)

type UserController struct {
	UserService *services.UserService
}

func NewUsuarioController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}
