package services

import (
	"e-commerce/internal/core/ports"
)

type DataBaseService struct {
	DataBaseRepository ports.DataBaseRepository
}

func NewDataBaseService(dataBaseRepository ports.DataBaseRepository) *DataBaseService {
	return &DataBaseService{
		DataBaseRepository: dataBaseRepository,
	}
}
