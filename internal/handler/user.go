package handler

import (
	"e-commerce/internal/core/domain"
	"e-commerce/internal/core/ports"
	"e-commerce/internal/error"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	commonHandler
	UserService ports.UserService
}

func NewUserHandler(userService ports.UserService) *UserHandler {
	return &UserHandler{
		UserService: userService,
	}
}
func (h UserHandler) GetPaginatedList(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	requestParams := h.GetRequestParams(ctx)
	page, errorParams := h.UserService.GetPaginatedList(requestParams)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(page)
}

func (h UserHandler) Get(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	var userName string
	id := ctx.Params("id", "")
	userName = ctx.Locals("username").(string)
	user, errorParams := h.UserService.Get(id, userName)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(user)
}

func (h UserHandler) Update(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	userUpdateRequest = new(domain.UserUpdateRequest)
	var userName string
	id := ctx.Params("id", "")
	userName = ctx.Locals("username").(string)
	err := ctx.BodyParser(userToUpdate)
	if err != nil {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Failed to parse request body",
			Status:  fiber.StatusBadRequest,
		})
	}
	errorParams := h.UserService.Update(id, userName, userToUpdate)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}
