package services

import (
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/error"
)

type ProductService struct {
	ProductRepository ports.ProductRepository
}

func NewProductService(productRepository ports.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}
func (s *ProductService) GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams) {
	page := new(models.Page)
	queryParams := page.SetRequestParamsAndGetQueryParams(requestParams)
	content, totalCount, errorParams := s.ProductRepository.FindPaginatedWithTotalCount(queryParams)
	if errorParams != nil {
		return nil, errorParams
	}
	page.SetResultParams(content, totalCount)
	return page, nil
}
func (s *ProductService) Get(id int) (*models.Product, *httpError.ErrorParams) {
	product, errorParams := s.ProductRepository.FindById(id)
	if errorParams != nil {
		return nil, errorParams
	}
	return product, nil
}
func (s *ProductService) Post(product *models.Product) (int, *httpError.ErrorParams) {
	id, err := s.ProductRepository.Insert(product)
	if err != nil {
		return 0, err
	}
	return id, nil
}
