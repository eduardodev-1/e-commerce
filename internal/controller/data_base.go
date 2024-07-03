package controller

import (
	"e-commerce/internal/database"
	"e-commerce/internal/services"
	"github.com/gofiber/fiber/v2"
)

type DataBaseController struct {
	DataBaseService *services.DataBaseService
}

func NewDataBaseController(DataBaseService *services.DataBaseService) *DataBaseController {
	return &DataBaseController{
		DataBaseService: DataBaseService,
	}
}
func (c *DataBaseController) ResetDataBase(ctx *fiber.Ctx) error {
	db := c.DataBaseService.DataBaseRepository.DB
	err := database.ExecuteSQLFile("internal/database/deleteSchema.sql", db)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error deleteSchema": err.Error(),
		})
	}
	err = database.ExecuteSQLFile("internal/database/schema.sql", db)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error execute schema": err.Error(),
		})
	}
	err = database.ExecuteSQLFile("internal/database/inserts.sql", db)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error inserts": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "restart database success.",
	})
}
