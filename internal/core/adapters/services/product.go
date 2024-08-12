package services

import (
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/httperror"
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
func (s *ProductService) Post(product *models.Product, username string, isAdmin bool) (int, *httpError.ErrorParams) {
	id, err := s.ProductRepository.Insert(product, username, isAdmin)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *ProductService) Update(product *models.Product, username string, isAdmin bool) *httpError.ErrorParams {
	errorParams := s.ProductRepository.Update(product, username, isAdmin)
	if errorParams != nil {
		return errorParams
	}
	return nil
}
func (s *ProductService) Delete(id int, username string, isAdmin bool) *httpError.ErrorParams {
	errorParams := s.ProductRepository.Delete(id, username, isAdmin)
	if errorParams != nil {
		return errorParams
	}
	return nil
}
