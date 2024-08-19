package fiber

import (
	"e-commerce/internal/auth"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/httperror"
	"github.com/go-oauth2/oauth2/v4"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/basicauth"
)

type AuthHandler struct {
	UserService ports.UserService
}

func NewAuthHandler(userService ports.UserService) *AuthHandler {
	return &AuthHandler{
		UserService: userService,
	}
}

func (h *AuthHandler) Authenticate(ctx *fiber.Ctx) error {
	customError := httpError.HttpCustomError{Ctx: ctx}
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
		errorParams := new(httpError.ErrorParams)
		response, errorParams = h.UserService.AuthenticateUserWithPasswordCredentials(authenticateRequest)
		if errorParams != nil {
			return customError.NewHttpError(errorParams)
		}
	default:
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid grantType",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (h *AuthHandler) Post(ctx *fiber.Ctx) error {
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
