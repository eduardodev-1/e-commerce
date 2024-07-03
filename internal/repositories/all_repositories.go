package repositories

import "github.com/jmoiron/sqlx"

type Repositories struct {
	UserRepository    UserRepository
	ProductRepository ProductRepository
}

func NewRepositories(db *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository:    NewUserRepository(db),
		ProductRepository: NewProductRepository(db),
	}
}
