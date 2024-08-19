package fiber

import (
	"e-commerce/internal/core/domain/models"
	"github.com/gofiber/fiber/v2"
)

type commonHandler struct{}

func (*commonHandler) GetRequestParams(ctx *fiber.Ctx) *models.RequestParams {
	requestParams := new(models.RequestParams)
	requestParams.PageNumber = ctx.QueryInt("page", 0)
	requestParams.PageSize = ctx.QueryInt("size", 10)
	requestParams.SortBy = ctx.Query("sort", "id")
	return requestParams
}
