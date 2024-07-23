package handlers

import (
	"e-commerce/internal/core/domain/models"
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
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	user, errorParams := h.UserService.Get(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(user)
}

func (h UserHandler) Update(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	userUpdateRequest := new(models.UserUpdateRequest)
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	err = ctx.BodyParser(userUpdateRequest)
	if err != nil {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Failed to parse request body",
			Status:  fiber.StatusBadRequest,
		})
	}
	userUpdateRequest.User.ID = id
	errorParams := h.UserService.Update(userUpdateRequest)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
	})
}

func (h UserHandler) Delete(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	errorParams := h.UserService.Delete(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
