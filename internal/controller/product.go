package controller

import (
	"e-commerce/internal/models"
	"e-commerce/internal/services"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	commonController
	ProductService services.ProductService
}

func (c ProductController) GetPaginatedList(ctx *fiber.Ctx) error {
	var page = new(models.Page)
	requestParams := GetRequestParams(ctx)
	page, err := c.ProductService.GetPaginatedList(requestParams)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(page)
}

func NewProductController(ProductSvc services.ProductService) ProductController {
	return ProductController{
		ProductService: ProductSvc,
	}
}
