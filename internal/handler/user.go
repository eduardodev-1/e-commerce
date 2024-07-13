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
func (c UserHandler) GetPaginatedList(ctx *fiber.Ctx) error {

	fiberError := http_error.HttpCustomError{Ctx: ctx}
	var page = new(domain.Page)
	requestParams := c.GetRequestParams(ctx)
	page, errorParams := c.UserService.GetPaginatedList(requestParams, page)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(page)
}

func (c UserHandler) Get(ctx *fiber.Ctx) error {
	fiberError := http_error.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		errorParams := new(http_error.ErrorParams)
		errorParams.Message = "invalid id"
		errorParams.Status = fiber.StatusBadRequest
		return fiberError.NewHttpError(errorParams)
	}
	product, errorParams := c.UserService.Get(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(product)
}
