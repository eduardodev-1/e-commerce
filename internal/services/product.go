package services

import (
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
)

type ProductService struct {
	ProductRepository repositories.ProductRepository
}

func NewProductService(productRepository repositories.ProductRepository) ProductService {
	return ProductService{
		ProductRepository: productRepository,
	}
}

func (s ProductService) GetPaginatedList(requestParams *models.RequestParams) (*models.Page, error) {
	var err error
	page := new(models.Page)
	page.SetRequestParams(requestParams)
	queryParams := page.GetQueryParams()
	content, count, err := s.ProductRepository.PageableFindAll(queryParams)
	if err != nil {
		return nil, err
	}
	page.SetResultParams(content, count)
	return page, nil
}
