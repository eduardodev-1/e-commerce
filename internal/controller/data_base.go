package controller

import (
	"e-commerce/internal/core/services"
	"e-commerce/internal/database/postgres"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
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
	db := c.DataBaseService.DataBaseRepository.GetDB().(*sqlx.DB)
	err := postgres.ExecuteSQLFile("internal/database/postgres/deleteSchema.sql", db)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error deleteSchema": err.Error(),
		})
	}
	err = postgres.ExecuteSQLFile("internal/database/postgres/schema.sql", db)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error execute schema": err.Error(),
		})
	}
	err = postgres.ExecuteSQLFile("internal/database/postgres/inserts.sql", db)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error inserts": err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": "restart database success.",
	})
}
