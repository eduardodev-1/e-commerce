package repositories

import (
	"database/sql"
	"e-commerce/internal/core/domain"
	"e-commerce/internal/error"
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
func (r *UserRepository) GetAuthenticationData(username string) (*domain.AuthenticatedUser, string, error) {
	var result struct {
		domain.AuthenticatedUser
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
func (r *UserRepository) FindPaginatedWithTotalCount(params *domain.QueryParams) (*[]domain.User, domain.TotalCount, *http_error.ErrorParams) {
	var users = new([]domain.User)
	var errorParams = new(http_error.ErrorParams)
	var total domain.TotalCount

	countQuery := `SELECT COUNT(*) FROM tb_user`
	err := r.Get(&total, countQuery)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, total, errorParams
	}
	query := fmt.Sprintf(`SELECT * FROM tb_user ORDER BY %s LIMIT $1 OFFSET $2`, params.Order)
	err = r.Select(users, query, params.Limit, params.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return users, total, nil
		}
		errorParams.SetDefaultParams(err)
		return nil, total, errorParams
	}
	return users, total, nil
}

func (r *UserRepository) GetById(id int) (*domain.User, *http_error.ErrorParams) {
	var errorParams = new(http_error.ErrorParams)
	var user = new(domain.User)
	query := `SELECT * FROM tb_user WHERE id = $1`
	err := r.Get(user, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorParams.SetCustomError(fiber.StatusNotFound, "User not found")
			return nil, errorParams
		}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	return user, errorParams
}
func (r *UserRepository) Insert(user *domain.User) (*domain.User, *http_error.ErrorParams) {
	var errorParams = new(http_error.ErrorParams)
	tx, err := r.Beginx()
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}

	// Inserir localização primeiro
	var locationID int
	locationQuery := `
		INSERT INTO tb_location (
			street_number, street_name, city, state, country, postcode, coordinates_latitude, 
			coordinates_longitude, timezone_offset, timezone_description
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	err = tx.QueryRowx(locationQuery,
		user.Location.Street.Number, user.Location.Street.Name, user.Location.City, user.Location.State,
		user.Location.Country, user.Location.Postcode, user.Location.Coordinates.Latitude,
		user.Location.Coordinates.Longitude, user.Location.Timezone.Offset, user.Location.Timezone.Description).Scan(&locationID)
	if err != nil {
		tx.Rollback()
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}

	// Inserir usuário
	var userID int
	userQuery := `
		INSERT INTO tb_user (
			gender, name_title, name_first, name_last, email, login_uuid, login_username, 
			login_password, login_salt, login_md5, login_sha1, login_sha256, dob_date, dob_age, 
			registered_date, registered_age, phone, cell, id_name, id_value, picture_large, 
			picture_medium, picture_thumbnail, nat, location_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, 
			$19, $20, $21, $22, $23, $24, $25) RETURNING id`
	err = tx.QueryRowx(userQuery,
		user.Gender, user.Name.Title, user.Name.First, user.Name.Last, user.Email, user.Login.UUID,
		user.Login.Username, user.Login.Password, user.Login.Salt, user.Login.MD5, user.Login.SHA1,
		user.Login.SHA256, user.Dob.Date, user.Dob.Age, user.Registered.Date, user.Registered.Age,
		user.Phone, user.Cell, user.Id.Name, user.Id.Value, user.Picture.Large, user.Picture.Medium,
		user.Picture.Thumbnail, user.Nat, locationID).Scan(&userID)
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			errorParams.SetDefaultParams(rollBackErr)
			return nil, errorParams
		}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}

	err = tx.Commit()
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	user.ID = userID // Set the returned user ID
	return user, errorParams
}
