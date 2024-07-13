package handler

import (
	"e-commerce/internal/core/domain"
	"e-commerce/internal/core/ports"
	"e-commerce/internal/error"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	commonHandler
	ProductService ports.ProductService
}

func NewProductHandler(ProductSvc ports.ProductService) *ProductHandler {
	return &ProductHandler{
		ProductService: ProductSvc,
	}
}
func (c ProductHandler) GetPaginatedList(ctx *fiber.Ctx) error {
	httpCustomError := http_error.HttpCustomError{Ctx: ctx}
	var page = new(domain.Page)
	requestParams := c.GetRequestParams(ctx)
	page, errorParams := c.ProductService.GetPaginatedList(requestParams, page)
	if errorParams != nil {
		return httpCustomError.NewHttpError(errorParams)
	}
	return ctx.JSON(page)
}

func (c ProductHandler) Get(ctx *fiber.Ctx) error {
	fiberError := http_error.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		errorParams := new(http_error.ErrorParams)
		errorParams.Message = "invalid id"
		errorParams.Status = fiber.StatusBadRequest
		return fiberError.NewHttpError(errorParams)
	}
	product, errorParams := c.ProductService.Get(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(product)
}
