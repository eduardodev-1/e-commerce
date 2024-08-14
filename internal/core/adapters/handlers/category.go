package handlers

import (
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	"e-commerce/internal/httperror"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	commonHandler
	CategoryService ports.CategoryService
}

func NewCategoryHandler(CategorySvc ports.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		CategoryService: CategorySvc,
	}
}
func (h CategoryHandler) GetPaginatedList(ctx *fiber.Ctx) error {
	httpCustomError := httpError.HttpCustomError{Ctx: ctx}
	requestParams := h.GetRequestParams(ctx)
	page, errorParams := h.CategoryService.GetPaginatedList(requestParams)
	if errorParams != nil {
		return httpCustomError.NewHttpError(errorParams)
	}
	return ctx.JSON(page)
}

func (h CategoryHandler) Get(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	category, errorParams := h.CategoryService.Get(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.JSON(category)
}

func (h CategoryHandler) Post(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	var category = new(models.Category)
	if err := ctx.BodyParser(&category); err != nil {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Failed to parse request body",
			Status:  fiber.StatusBadRequest,
		})
	}
	id, errorParams := h.CategoryService.Post(category)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}

	// Retornar resposta de sucesso
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "category created successfully",
		"id":      id,
	})
}

func (h CategoryHandler) Update(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	var category = new(models.Category)
	if err = ctx.BodyParser(&category); err != nil {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "Failed to parse request body",
			Status:  fiber.StatusBadRequest,
		})
	}
	category.ID = int64(id)
	errorParams := h.CategoryService.Update(category)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}

	// Retornar resposta de sucesso
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "category updated successfully",
		"id":      id,
	})
}

func (h CategoryHandler) Delete(ctx *fiber.Ctx) error {
	fiberError := httpError.HttpCustomError{Ctx: ctx}
	id, err := ctx.ParamsInt("id", 0)
	if err != nil || id <= 0 {
		return fiberError.NewHttpError(&httpError.ErrorParams{
			Message: "invalid id",
			Status:  fiber.StatusBadRequest,
		})
	}
	errorParams := h.CategoryService.Delete(id)
	if errorParams != nil {
		return fiberError.NewHttpError(errorParams)
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}
