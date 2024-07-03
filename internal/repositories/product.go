package repositories

import (
	"database/sql"
	"e-commerce/internal/models"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	*sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db,
	}
}
func (r ProductRepository) PageableFindAll(params *models.QueryParams) (*[]models.Product, int, error) {
	var product = new([]models.Product)
	var total int

	countQuery := `SELECT COUNT(*) FROM tb_product`
	err := r.Get(&total, countQuery)
	if err != nil {
		return nil, total, err
	}

	query := fmt.Sprintf(`SELECT * FROM tb_product ORDER BY %s LIMIT $1 OFFSET $2`, params.Order)
	err = r.Select(product, query, params.Limit, params.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return product, total, err
		}
		return nil, total, err
	}
	return product, total, nil
}
