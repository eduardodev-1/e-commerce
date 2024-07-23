package handlers

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/error"
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

func (h *LoginHandler) Authenticate(ctx *fiber.Ctx) error {
	err := auth.CheckAppCredentials(ctx)
	if err != nil {
		return err
	}
	authenticateRequest := new(models.RequestCredentials)
	if err = ctx.BodyParser(authenticateRequest); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	var response interface{}
	// Select authentication according to grand type.
	grantType := oauth2.GrantType(authenticateRequest.GrantType)
	switch grantType {
	case oauth2.PasswordCredentials:
		response, err = h.UserService.AuthenticateUserWithPasswordCredentials(authenticateRequest)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid grantType",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *LoginHandler) Post(ctx *fiber.Ctx) error {
	newUser := new(models.UserFromRequest)
	customError := httpError.HttpCustomError{Ctx: ctx}
	errorParams := new(httpError.ErrorParams)
	if err := ctx.BodyParser(newUser); err != nil {
		errorParams.SetDefaultParams(err)
		return customError.NewHttpError(errorParams)
	}
	_, errorParams = h.UserService.CreateNewUser(newUser)
	if errorParams != nil {
		return customError.NewHttpError(errorParams)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
