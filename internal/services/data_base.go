package services

import (
	"e-commerce/internal/repositories"
)

type DataBaseService struct {
	DataBaseRepository *repositories.DataBaseRepository
}

func NewDataBaseService(productRepository *repositories.DataBaseRepository) *DataBaseService {
	return &DataBaseService{
		DataBaseRepository: productRepository,
	}
}
