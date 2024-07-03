package controller

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"fmt"
	"github.com/go-oauth2/oauth2/v4"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/basicauth"
)

type LoginController struct {
	UserService *services.UserService
}

func NewLoginController(userService *services.UserService) *LoginController {
	return &LoginController{
		UserService: userService,
	}
}

func (h *LoginController) Autenticate(c *fiber.Ctx) error {
	err := auth.CheckAppCredentials(c)
	if err != nil {
		return err
	}

	loginRequest := new(models.RequestCredentials)
	if err = c.BodyParser(loginRequest); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var user = new(models.AuthenticatedUser)
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
