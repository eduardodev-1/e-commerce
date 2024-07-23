package models

import (
	"math"
)

type Page struct {
	Content       interface{} `json:"content"`
	TotalElements TotalCount  `json:"totalElements"`
	TotalPages    int         `json:"totalPages"`
	Size          int         `json:"size"`
	Number        int         `json:"number"`
	Sort          string      `json:"sort"`
}
type RequestParams struct {
	PageNumber int    `query:"page"`
	PageSize   int    `query:"size"`
	SortBy     string `query:"sort"`
}
type QueryParams struct {
	Limit  int
	Offset int
	Order  string
}
type TotalCount int

func (p *Page) SetRequestParamsAndGetQueryParams(params *RequestParams) *QueryParams {
	p.SetRequestParams(params)
	return p.GetQueryParams()
}
func (p *Page) SetRequestParams(params *RequestParams) {
	p.Number = params.PageNumber
	p.Size = params.PageSize
	p.Sort = params.SortBy
}
func (p *Page) GetQueryParams() *QueryParams {
	return &QueryParams{
		Limit:  p.Size,
		Offset: p.Number * p.Size,
		Order:  p.Sort,
	}
}
func (p *Page) SetResultParams(content interface{}, count TotalCount) {
	p.Content = content
	p.TotalElements = count
	p.TotalPages = int(math.Ceil(float64(p.TotalElements) / float64(p.Size)))
}
