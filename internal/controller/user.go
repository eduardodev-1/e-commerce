package controller

import (
	"e-commerce/internal/core/domain"
	"e-commerce/internal/core/services"
	"e-commerce/internal/error"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	commonController
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}
func (c UserController) GetPaginatedList(ctx *fiber.Ctx) error {

	fiberError := http_error.HttpCustomError{Ctx: ctx}
	var page = new(domain.Page)
	requestParams := c.GetRequestParams(ctx)
	page, errorParams := c.UserService.GetPaginatedList(requestParams, page)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(page)
}

func (c UserController) Get(ctx *fiber.Ctx) error {
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
