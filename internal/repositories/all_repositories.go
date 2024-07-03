package repositories

import "github.com/jmoiron/sqlx"

type Repositories struct {
	UsuarioRepository UsuarioRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UsuarioRepository: NewUsuarioRepository(db),
	}
}
