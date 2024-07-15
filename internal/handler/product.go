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
func (h ProductHandler) GetPaginatedList(ctx *fiber.Ctx) error {
	httpCustomError := httpError.HttpCustomError{Ctx: ctx}
	requestParams := h.GetRequestParams(ctx)
	page, errorParams := h.ProductService.GetPaginatedList(requestParams)
	if errorParams != nil {
		return httpCustomError.NewHttpError(errorParams)
	}
	return ctx.JSON(page)
}

func (h ProductHandler) Get(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		errorParams := new(httpError.ErrorParams)
		errorParams.Message = "invalid id"
		errorParams.Status = fiber.StatusBadRequest
		return fiberError.NewHttpError(errorParams)
	}
	product, errorParams := h.ProductService.Get(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(product)
}

func (h ProductHandler) Post(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	var product = new(domain.Product)
	if err := ctx.BodyParser(&product); err != nil {
		errorParams := new(httpError.ErrorParams)
		errorParams.Message = "Failed to parse request body"
		errorParams.Status = fiber.StatusBadRequest
		return fiberError.NewHttpError(errorParams)
	}
	id, errorParams := h.ProductService.Post(product)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}

	// Retornar resposta de sucesso
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "product created successfully",
		"id":      id,
	})
}
