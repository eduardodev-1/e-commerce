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

type CategoryRepository struct {
	*sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{
		db,
	}
}
func (r *CategoryRepository) FindPaginatedWithTotalCount(params *models.QueryParams) (*[]models.Category, models.TotalCount, *fiber_error.ErrorParams) {

	var products = new([]models.Category)
	var errorParams = new(fiber_error.ErrorParams)
	var total models.TotalCount

	countQuery := `SELECT COUNT(*) FROM tb_category`
	err := r.Get(&total, countQuery)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, total, errorParams
	}
	query := fmt.Sprintf(`SELECT * FROM tb_category ORDER BY %s LIMIT $1 OFFSET $2`, params.Order)
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

func (r *CategoryRepository) FindById(id int) (*models.Category, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	var product = new(models.Category)
	query := `SELECT * FROM tb_category WHERE id = $1`
	err := r.Get(product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorParams.SetCustomError(fiber.StatusNotFound, "Category not found")
			return nil, errorParams
		}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	return product, nil
}
func (r *CategoryRepository) Insert(category *models.Category) (int, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	query := `INSERT INTO tb_category (name)
	VALUES ($1) RETURNING id`
	var id int
	err := r.QueryRow(query, category.Name).Scan(&id)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return id, errorParams
	}
	return id, nil
}

func (r *CategoryRepository) Update(category *models.Category) *fiber_error.ErrorParams {
	var errorParams = new(fiber_error.ErrorParams)
	var result sql.Result
	var err error
	var rowsAffected int64
	updateQuery := `UPDATE tb_category SET name = $1 WHERE id = $2`
	result, err = r.Exec(updateQuery, category.Name, category.ID)
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
		errorParams.SetCustomError(fiber.StatusNotFound, "Category not found")
		return errorParams
	}
	return nil
}
func (r *CategoryRepository) Delete(id int) *fiber_error.ErrorParams {
	var errorParams = new(fiber_error.ErrorParams)
	deleteQuery := `DELETE FROM tb_category WHERE id = $1`
	result, err := r.Exec(deleteQuery, id)
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
		errorParams.SetDefaultParams(errors.New("nenhuma categoria encontrada com este ID"))
		return errorParams
	}
	return nil
}
