package controller

import (
	fiber_error "e-commerce/internal/error"
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	ProductService *services.ProductService
}

func NewProductController(ProductSvc *services.ProductService) *ProductController {
	return &ProductController{
		ProductService: ProductSvc,
	}
}
func (c ProductController) GetPaginatedList(ctx *fiber.Ctx) error {
	var page = new(models.Page)
	requestParams := page.GetRequestParams(ctx)
	page, err := c.ProductService.GetPaginatedList(requestParams, page)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(page)
}

func (c ProductController) Get(ctx *fiber.Ctx) error {
	fiberError := fiber_error.FiberCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		errorParams := new(fiber_error.ErrorParams)
		errorParams.Message = "invalid id"
		errorParams.Status = fiber.StatusBadRequest
		return fiberError.NewFiberError(errorParams)
	}
	product, errorParams := c.ProductService.Get(id)
	if errorParams != nil {
		return fiberError.NewFiberError(errorParams)
	}
	return ctx.JSON(product)
}
