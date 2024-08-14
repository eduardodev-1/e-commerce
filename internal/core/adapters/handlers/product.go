package handlers

import (
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	"e-commerce/internal/httperror"
	"e-commerce/internal/network/middleware"
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
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	product, errorParams := h.ProductService.Get(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(product)
}

func (h ProductHandler) Post(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	username := ctx.Locals("username").(string)
	authorities := ctx.Locals("authorities").([]string)
	var product = new(models.Product)
	if err := ctx.BodyParser(&product); err != nil {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Failed to parse request body",
			Status:  fiber.StatusBadRequest,
		})
	}
	isAdmin := middleware.HasRole(authorities, "ROLE_ADMIN")
	id, errorParams := h.ProductService.Post(product, username, isAdmin)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}

	// Retornar resposta de sucesso
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "product created successfully",
		"id":      id,
	})
}

func (h ProductHandler) Update(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	username := ctx.Locals("username").(string)
	var product = new(models.Product)
	if err := ctx.BodyParser(&product); err != nil {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Failed to parse request body",
			Status:  fiber.StatusBadRequest,
		})
	}
	product.ID = int64(id)
	authorities := ctx.Locals("authorities").([]string)
	isAdmin := middleware.HasRole(authorities, "ROLE_ADMIN")
	errorParams := h.ProductService.Update(product, username, isAdmin)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}

	// Retornar resposta de sucesso
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "product updated successfully",
		"id":      id,
	})
}

func (h ProductHandler) Delete(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	username := ctx.Locals("username").(string)
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	authorities := ctx.Locals("authorities").([]string)
	isAdmin := middleware.HasRole(authorities, "ROLE_ADMIN")
	errorParams := h.ProductService.Delete(id, username, isAdmin)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
