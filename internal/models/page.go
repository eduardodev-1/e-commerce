package models

import (
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
	Number int
	Size   int
	Sort   string
}
type QueryParams struct {
	Limit  int
	Offset int
	Order  string
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
