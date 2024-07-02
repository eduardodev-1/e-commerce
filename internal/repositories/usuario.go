package repositories

import (
	"fmt"
	"github.com/eduardodev-1/e-commerce/internal/models"
	"github.com/jmoiron/sqlx"
)

type UsuarioRepository struct {
	*sqlx.DB
}

func NewUsuarioRepository(db *sqlx.DB) UsuarioRepository {
	return UsuarioRepository{
		db,
	}
}
func (r *UsuarioRepository) GetAuthoritiesByUserName(username string) ([]string, error) {
	query := `SELECT authority
				FROM tb_role
				INNER JOIN tb_user_role ON tb_role.id = tb_user_role.role_id
				INNER JOIN tb_user ON tb_user.id = tb_user_role.user_id
				WHERE tb_user.email = $1`
	rows, err := r.Query(query, username)
	if err != nil {
		return nil, err
	}

	var authorities []string

	for rows.Next() {
		var authority string

		err = rows.Scan(&authority)
		// TODO: avaliar este erro eh necessario?
		if err != nil {
			fmt.Println(err)
		}

		authorities = append(authorities, authority)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return authorities, nil
}
func (r *UsuarioRepository) GetAuthenticationData(username string) (*models.AuthenticatedUser, string, error) {
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
		return nil, "", err
	}

	u := &result.AuthenticatedUser
	return u, result.HashedPassword, nil
}
