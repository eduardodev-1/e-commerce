package controller

import (
	"e-commerce/internal/models"
	"github.com/gofiber/fiber/v2"
)

type commonController struct{}

func GetRequestParams(ctx *fiber.Ctx) *models.RequestParams {
	requestParams := new(models.RequestParams)
	requestParams.Number = ctx.QueryInt("page", 0)
	requestParams.Size = ctx.QueryInt("size", 10)
	requestParams.Sort = ctx.Query("sort", "id")
	return requestParams
}
