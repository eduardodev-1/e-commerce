package repositories

import (
	"database/sql"
	"e-commerce/internal/models"
	"errors"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	*sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepository{
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
