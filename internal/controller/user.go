package controller

import (
	fiber_error "e-commerce/internal/error"
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		UserService: userService,
	}
}
func (c UserController) GetPaginatedList(ctx *fiber.Ctx) error {

	fiberError := fiber_error.FiberCustomError{Ctx: ctx}
	var page = new(models.Page)
	requestParams := page.GetRequestParams(ctx)
	page, errorParams := c.UserService.GetPaginatedList(requestParams, page)
	if errorParams != nil {
		return fiberError.NewFiberError(errorParams)
	}
	return ctx.JSON(page)
}

func (c UserController) Get(ctx *fiber.Ctx) error {
	fiberError := fiber_error.FiberCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		errorParams := new(fiber_error.ErrorParams)
		errorParams.Message = "invalid id"
		errorParams.Status = fiber.StatusBadRequest
		return fiberError.NewFiberError(errorParams)
	}
	product, errorParams := c.UserService.Get(id)
	if errorParams != nil {
		return fiberError.NewFiberError(errorParams)
	}
	return ctx.JSON(product)
}
