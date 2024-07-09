package services

import (
	fiber_error "e-commerce/internal/error"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type ProductService struct {
	ProductRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s ProductService) GetPaginatedList(requestParams *models.RequestParams, page *models.Page) (*models.Page, *fiber_error.ErrorParams) {
	page.SetRequestParams(requestParams)
	queryParams := page.GetQueryParams()
	content, count, errorParams := s.ProductRepository.PageableFindAll(queryParams)
	if errorParams != nil {
		return nil, errorParams
	}
	page.SetResultParams(content, count)
	return page, nil
}

func (s ProductService) Get(id int) (*models.Product, *fiber_error.ErrorParams) {
	product, errorParams := s.ProductRepository.GetById(id)
	if errorParams != nil {
		return nil, errorParams
	}
	return product, nil
}
