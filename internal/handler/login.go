package handler

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/core/domain"
	"e-commerce/internal/core/ports"
	"fmt"
	"github.com/go-oauth2/oauth2/v4"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/basicauth"
)

type LoginHandler struct {
	UserService ports.UserService
}

func NewLoginHandler(userService ports.UserService) *LoginHandler {
	return &LoginHandler{
		UserService: userService,
	}
}

func (h *LoginHandler) Authenticate(c *fiber.Ctx) error {
	err := auth.CheckAppCredentials(c)
	if err != nil {
		return err
	}

	loginRequest := new(domain.RequestCredentials)
	if err = c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user = new(domain.AuthenticatedUser)
	grantType := oauth2.GrantType(loginRequest.GrantType)

	switch grantType {
	case oauth2.PasswordCredentials:
		user, err = h.UserService.Authenticate(loginRequest)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	default:
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid grantType",
		})
	}
	roles, err := h.UserService.GetUserRoles(user.Username)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	fmt.Println(roles)
	// Criar JWToken
	token, err := auth.NewJWToken(user.Id, user.Username, roles)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.LoginResponse{Token: token})
}
