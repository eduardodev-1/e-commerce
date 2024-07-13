package services

import (
	"e-commerce/internal/core/domain"
	"e-commerce/internal/core/ports"
	fiber_error "e-commerce/internal/error"
)

type ProductService struct {
	ProductRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s ProductService) GetPaginatedList(requestParams *domain.RequestParams, page *domain.Page) (*domain.Page, *fiber_error.ErrorParams) {
	queryParams := page.SetRequestParamsAndGetQueryParams(requestParams)
	content, totalCount, errorParams := s.ProductRepository.FindPaginatedWithTotalCount(queryParams)
	if errorParams != nil {
		return nil, errorParams
	}
	page.SetResultParams(content, totalCount)
	return page, nil
}

func (s ProductService) Get(id int) (*domain.Product, *fiber_error.ErrorParams) {
	product, errorParams := s.ProductRepository.GetById(id)
	if errorParams != nil {
		return nil, errorParams
	}
	return product, nil
}
