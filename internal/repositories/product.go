package repositories

import (
	"database/sql"
	"e-commerce/internal/models"
	"errors"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	*sqlx.DB
}

func (r ProductRepository) PageableFindAll(params *models.QueryParams) (*[]models.Product, int, error) {
	var product = new([]models.Product)
	var total int

	// Consulta para contar o número total de produtos
	countQuery := `SELECT COUNT(*) FROM tb_product`
	err := r.Get(&total, countQuery)
	if err != nil {
		return nil, total, err
	}

	query := `SELECT * FROM tb_product ORDER BY $1 LIMIT $2 OFFSET $3`
	err = r.Select(product, query, params.Order, params.Limit, params.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return product, total, err
		}
		return nil, total, err
	}
	return product, total, nil
}

func NewProductRepository(db *sqlx.DB) ProductRepository {
	return ProductRepository{
		db,
	}
}
