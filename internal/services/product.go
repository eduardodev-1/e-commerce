package services

import (
	fiber_error "e-commerce/internal/error"
	"e-commerce/internal/models"
	"e-commerce/internal/repositories"
	"log"
)

type ProductService struct {
	ProductRepository *repositories.ProductRepository
}

func NewProductService(productRepository *repositories.ProductRepository) *ProductService {
	return &ProductService{
		ProductRepository: productRepository,
	}
}

func (s ProductService) GetPaginatedList(requestParams *models.RequestParams, page *models.Page) (*models.Page, error) {
	var err error
	page.SetRequestParams(requestParams)
	log.Print(page)
	queryParams := page.GetQueryParams()
	log.Print(queryParams)
	content, count, err := s.ProductRepository.PageableFindAll(queryParams)
	if err != nil {
		return nil, err
	}
	page.SetResultParams(content, count)
	return page, nil
}

func (s ProductService) Get(id int) (*models.Product, *fiber_error.ErrorParams) {
	product := s.ProductRepository.GetById(id)

	return nil, nil
}
