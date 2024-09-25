package postgres

import (
	"database/sql"
	"e-commerce/internal/core/domain/models"
	fiber_error "e-commerce/internal/httperror"
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
func (r *ProductRepository) FindPaginatedWithTotalCount(params *models.QueryParams) (*[]models.Product, int, *fiber_error.ErrorParams) {

	var products = new([]models.Product)
	var errorParams = new(fiber_error.ErrorParams)
	var total int

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

func (r *ProductRepository) FindById(id int) (*models.Product, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	var product = new(models.Product)
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
func (r *ProductRepository) Insert(product *models.Product, username string, isAdmin bool) (int, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	if !isAdmin {
		getSellerIdQuery := `SELECT id FROM tb_user WHERE login_username = $1`
		err := r.Get(&product.SellerID, getSellerIdQuery, username)
		if err != nil {
			errorParams.SetDefaultParams(err)
			return 0, errorParams
		}
	}
	query := `INSERT INTO tb_product (price, description, img_url, name, seller, quantity)
	VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	var id int
	err := r.QueryRow(query, product.Price, product.Description, product.ImgURL, product.Name, product.SellerID, product.Quantity).Scan(&id)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return id, errorParams
	}
	return id, nil
}

func (r *ProductRepository) Update(product *models.Product, username string, isAdmin bool) *fiber_error.ErrorParams {
	var errorParams = new(fiber_error.ErrorParams)
	var result sql.Result
	var err error
	var rowsAffected int64
	if !isAdmin {
		getSellerIdQuery := `SELECT id FROM tb_user WHERE login_username = $1`
		err = r.Get(&product.SellerID, getSellerIdQuery, username)
		if err != nil {
			errorParams.SetDefaultParams(err)
			return errorParams
		}
		updateQuery := `
		UPDATE tb_product
		SET price = $1, description = $2, img_url = $3, name = $4, quantity = $5
		WHERE id = $6 AND seller = $7`

		result, err = r.Exec(updateQuery, product.Price, product.Description, product.ImgURL, product.Name, product.Quantity, product.ID, product.SellerID)
		if err != nil {
			errorParams.SetDefaultParams(err)
			return errorParams
		}
	} else {
		updateQuery := `
		UPDATE tb_product
		SET price = $1, description = $2, img_url = $3, name = $4, quantity = $5, seller= $6
		WHERE id = $7`
		result, err = r.Exec(updateQuery, product.Price, product.Description, product.ImgURL, product.Name, product.Quantity, product.SellerID, product.ID)
	}
	if err != nil {
		errorParams.SetDefaultParams(err)
		return errorParams
	}
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		errorParams.SetDefaultParams(err)
		return errorParams
	}
	if rowsAffected <= 0 {
		errorParams.SetCustomError(fiber.StatusNotFound, "Product not found")
		return errorParams
	}
	return nil
}
func (r *ProductRepository) Delete(id int, username string, isAdmin bool) *fiber_error.ErrorParams {
	var errorParams = new(fiber_error.ErrorParams)
	var deleteQuery string
	var result sql.Result
	var err error
	if !isAdmin {
		var sellerID int
		idQuery := `SELECT id FROM tb_user WHERE login_username = $1`
		err = r.Get(&sellerID, idQuery, username)
		if err != nil {
			errorParams.SetDefaultParams(err)
			return errorParams
		}
		deleteQuery = `
		DELETE FROM tb_product
		WHERE id = $1 AND seller = $2`
		result, err = r.Exec(deleteQuery, id, sellerID)
	} else {
		deleteQuery = `
		DELETE FROM tb_product
		WHERE id = $1`
		result, err = r.Exec(deleteQuery, id)
	}
	if err != nil {
		errorParams.SetDefaultParams(err)
		return errorParams
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		errorParams.SetDefaultParams(err)
		return errorParams
	}
	if rowsAffected == 0 {
		errorParams.SetCustomError(fiber.StatusNotFound, "nenhum produto encontrado para este vendedor com este ID")
		return errorParams
	}

	return nil
}
