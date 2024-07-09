package repositories

import (
	"database/sql"
	fiber_error "e-commerce/internal/error"
	"e-commerce/internal/models"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}
func (r *UserRepository) GetAuthoritiesByUserName(username string) ([]string, error) {
	var authorities []string
	query := `SELECT authority
				FROM tb_role
				INNER JOIN tb_user_role ON tb_role.id = tb_user_role.role_id
				INNER JOIN tb_user ON tb_user.id = tb_user_role.user_id
				WHERE tb_user.email = $1`
	err := r.Select(&authorities, query, username)
	if err != nil {
		return nil, err
	}
	return authorities, nil
}
func (r *UserRepository) GetAuthenticationData(username string) (*models.AuthenticatedUser, string, error) {
	var result struct {
		models.AuthenticatedUser
		HashedPassword string `db:"password"`
	}
	query := `
	SELECT u.id, u.email as username, u.password
	FROM tb_user u
	WHERE u.email = $1`

	err := r.Get(&result, query, username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, "", nil
		}
		return nil, "", err
	}

	u := &result.AuthenticatedUser
	return u, result.HashedPassword, nil
}
func (r UserRepository) PageableFindAll(params *models.QueryParams) (*[]models.User, int, *fiber_error.ErrorParams) {
	var products = new([]models.User)
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

func (r UserRepository) GetById(id int) (*models.User, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	var product = new(models.User)
	query := `SELECT * FROM tb_product WHERE id = $1`
	err := r.Get(product, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorParams.SetCustomError(fiber.StatusNotFound, "User not found")
			return nil, errorParams
		}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	return product, errorParams
}
func (r UserRepository) Insert(product *models.User) (*models.User, *fiber_error.ErrorParams) {
	var errorParams = new(fiber_error.ErrorParams)
	query := `INSERT INTO tb_product (price, description, img_url, name, seller)
	VALUES ($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := r.QueryRow(query, product.Price, product.Description, product.ImgURL).Scan(&id)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	return product, errorParams
}
