package repositories

import (
	"database/sql"
	"e-commerce/internal/core/domain"
	fiber_error "e-commerce/internal/error"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
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
func (r ProductRepository) FindPaginatedWithTotalCount(params *domain.QueryParams) (*[]domain.Product, domain.TotalCount, *fiber_error.ErrorParams) {

	var products = new([]domain.Product)
	var errorParams = new(fiber_error.ErrorParams)
	var total domain.TotalCount

	countQuery := `SELECT COUNT(*) FROM tb_product`
	err := r.Get(&total, countQuery)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, total, errorParams
	}
	query := fmt.Sprintf(`SELECT * FROM tb_product ORDER BY %s LIMIT $1 OFFSET $2`, params.Order)
	err = r.Select(products, query, params.Limit, params.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return products, total, nil
		}
		errorParams.SetDefaultParams(err)
		return nil, total, errorParams
	}
	return products, total, nil
}

func (r ProductRepository) FindById(id int) (*domain.Product, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	var product = new(domain.Product)
	query := `SELECT * FROM tb_product WHERE id = $1`
	err := r.Get(product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorParams.SetCustomError(fiber.StatusNotFound, "Product not found")
			return nil, errorParams
		}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	return product, nil
}
func (r ProductRepository) Insert(product *domain.Product) (domain.IdToResponse, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	query := `INSERT INTO tb_product (price, description, img_url, name, seller, quantity)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id domain.IdToResponse
	err := r.QueryRow(query, product.Price, product.Description, product.ImgURL).Scan(&id)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return id, errorParams
	}
	return id, errorParams
}
