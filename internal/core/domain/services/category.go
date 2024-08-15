package services

import (
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	httpError "e-commerce/internal/httperror"
)

type CategoryService struct {
	CategoryRepository ports.CategoryRepository
}

func NewCategoryService(productRepository ports.CategoryRepository) *CategoryService {
	return &CategoryService{
		CategoryRepository: productRepository,
	}
}
func (s *CategoryService) GetPaginatedList(requestParams *models.RequestParams) (*models.Page, *httpError.ErrorParams) {
	page := new(models.Page)
	queryParams := page.SetRequestParamsAndGetQueryParams(requestParams)
	content, totalCount, errorParams := s.CategoryRepository.FindPaginatedWithTotalCount(queryParams)
	if errorParams != nil {
		return nil, errorParams
	}
	page.SetResultParams(content, totalCount)
	return page, nil
}
func (s *CategoryService) Get(id int) (*models.Category, *httpError.ErrorParams) {
	product, errorParams := s.CategoryRepository.FindById(id)
	if errorParams != nil {
		return nil, errorParams
	}
	return product, nil
}
func (s *CategoryService) Post(category *models.Category) (int, *httpError.ErrorParams) {
	id, err := s.CategoryRepository.Insert(category)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *CategoryService) Update(category *models.Category) *httpError.ErrorParams {
	errorParams := s.CategoryRepository.Update(category)
	if errorParams != nil {
		return errorParams
	}
	return nil
}
func (s *CategoryService) Delete(id int) *httpError.ErrorParams {
	errorParams := s.CategoryRepository.Delete(id)
	if errorParams != nil {
		return errorParams
	}
	return nil
}
