package repositories

import (
	"database/sql"
	"e-commerce/internal/core/domain"
	"e-commerce/internal/error"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"log"
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
				WHERE tb_user.login_username = $1`
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
	SELECT u.id, u.login_username as username, u.login_password as password
	FROM tb_user u
	WHERE u.login_username = $1`

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
func (r *UserRepository) FindPaginatedWithTotalCount(params *domain.QueryParams) (*[]domain.User, domain.TotalCount, *httpError.ErrorParams) {
	var usersDB []domain.UserDB
	var users []domain.User
	var errorParams = new(httpError.ErrorParams)
	var total domain.TotalCount

	countQuery := `SELECT COUNT(*) FROM tb_user`
	err := r.Get(&total, countQuery)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, total, errorParams
	}
	query := fmt.Sprintf(`SELECT u.*, street_number, street_name, city, state, country, postcode, coordinates_latitude,
       coordinates_longitude, timezone_offset, timezone_description
				FROM
                    tb_user u
                JOIN
                        tb_location l ON u.location_id = l.id
				ORDER BY %s LIMIT $1 OFFSET $2`, params.Order)
	err = r.Select(&usersDB, query, params.Limit, params.Offset)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, total, nil
		}
		errorParams.SetDefaultParams(err)
		return nil, total, errorParams
	}
	for _, userDB := range usersDB {
		var user domain.User
		err := user.NewUserByUserDB(&userDB)
		if err != nil {
			return nil, 0, nil
		}
		users = append(users, user)
	}
	return &users, total, nil
}
func (r *UserRepository) FindByUserName(userName string) (*domain.User, *httpError.ErrorParams) {
	var errorParams = new(httpError.ErrorParams)
	var user = new(domain.User)
	var userDB = new(domain.UserDB)
	query := `SELECT u.*, street_number, street_name, city, state, country, postcode, coordinates_latitude,
       coordinates_longitude, timezone_offset, timezone_description
				FROM
                    tb_user u
                JOIN
                        tb_location l ON u.location_id = l.id
                WHERE u.login_username = $1`
	err := r.Get(userDB, query, userName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorParams.SetCustomError(fiber.StatusNotFound, "User not found")
			return nil, errorParams
		}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	err = user.NewUserByUserDB(userDB)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	return user, nil
}
func (r *UserRepository) FindById(id int) (*domain.User, *httpError.ErrorParams) {
	var errorParams = new(httpError.ErrorParams)
	var user = new(domain.User)
	var userDB = new(domain.UserDB)
	query := `SELECT u.*, street_number, street_name, city, state, country, postcode, coordinates_latitude,
       coordinates_longitude, timezone_offset, timezone_description
				FROM
                    tb_user u
                JOIN
                        tb_location l ON u.location_id = l.id
                WHERE u.id = $1`
	err := r.Get(userDB, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			errorParams.SetCustomError(fiber.StatusNotFound, "User not found")
			return nil, errorParams
		}
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	err = user.NewUserByUserDB(userDB)
	if err != nil {
		errorParams.SetDefaultParams(err)
		return nil, errorParams
	}
	return user, nil
}
func (r *UserRepository) Insert(newUser *domain.NewUserRequest) (int, *httpError.ErrorParams) {
	var errorParams = new(httpError.ErrorParams)
	var user = newUser.User
	tx, err := r.Beginx()
	if err != nil {
		errorParams.SetDefaultParams(err)
		return 0, errorParams
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
		errRollBack := tx.Rollback()
		if errRollBack != nil {
			errorParams.SetDefaultParams(errRollBack)
			return 0, errorParams
		}
		errorParams.SetDefaultParams(err)
		return 0, errorParams
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
			return 0, errorParams
		}
		errorParams.SetDefaultParams(err)
		return 0, errorParams
	}
	getAuthorityIdQuery := `
		SELECT id 
		FROM tb_role 
		WHERE authority ILIKE '%' || $1 || '%';
`
	var roleID = new(int)
	err = tx.Get(roleID, getAuthorityIdQuery, newUser.UserType)
	if err != nil {
		log.Print("erro ao tentar capturar o role id")
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			errorParams.SetDefaultParams(rollBackErr)
			return 0, errorParams
		}
		errorParams.SetDefaultParams(err)
		return 0, errorParams
	}

	insertAuthorityQuery := `
		INSERT INTO tb_user_role (role_id, user_id)
		VALUES ($1, $2)`

	err = tx.QueryRowx(insertAuthorityQuery, roleID, userID).Err()
	if err != nil {
		rollBackErr := tx.Rollback()
		if rollBackErr != nil {
			errorParams.SetDefaultParams(rollBackErr)
			return 0, errorParams
		}
		errorParams.SetDefaultParams(err)
		return 0, errorParams
	}
	err = tx.Commit()
	if err != nil {
		errorParams.SetDefaultParams(err)
		return 0, errorParams
	}
	return userID, nil
}
