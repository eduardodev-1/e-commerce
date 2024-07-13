package controller

import (
	"e-commerce/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type commonController struct{}

func (*commonController) GetRequestParams(ctx *fiber.Ctx) *domain.RequestParams {
	requestParams := new(domain.RequestParams)
	requestParams.PageNumber = ctx.QueryInt("page", 0)
	requestParams.PageSize = ctx.QueryInt("size", 10)
	requestParams.SortBy = ctx.Query("sort", "id")
	return requestParams
}
