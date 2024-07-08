package models

import (
	"github.com/gofiber/fiber/v2"
	"math"
)

type Page struct {
	Content       interface{} `json:"content"`
	TotalElements int         `json:"totalElements"`
	TotalPages    int         `json:"totalPages"`
	Size          int         `json:"size"`
	Number        int         `json:"number"`
	Sort          string      `json:"sort"`
}
type RequestParams struct {
	Number int    `query:"page"`
	Size   int    `query:"size"`
	Sort   string `query:"sort"`
}
type QueryParams struct {
	Limit  int
	Offset int
	Order  string
}

func (p *Page) GetRequestParams(ctx *fiber.Ctx) *RequestParams {
	requestParams := new(RequestParams)
	requestParams.Number = ctx.QueryInt("page", 0)
	requestParams.Size = ctx.QueryInt("size", 10)
	requestParams.Sort = ctx.Query("sort", "id")
	return requestParams
}

func (p *Page) GetQueryParams() *QueryParams {
	return &QueryParams{
		Limit:  p.Size,
		Offset: p.Number * p.Size,
		Order:  p.Sort,
	}
}

func (p *Page) SetRequestParams(params *RequestParams) {
	p.Number = params.Number
	p.Size = params.Size
	p.Sort = params.Sort
}

func (p *Page) SetResultParams(content interface{}, count int) {
	p.Content = content
	p.TotalElements = count
	p.TotalPages = int(math.Ceil(float64(p.TotalElements) / float64(p.Size)))
}
