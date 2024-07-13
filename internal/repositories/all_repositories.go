package repositories

import (
	"e-commerce/internal/core/ports"
	"errors"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UserRepository     ports.UserRepository
	ProductRepository  ports.ProductRepository
	DataBaseRepository ports.DataBaseRepository
}

func NewRepositories(db interface{}, dbType string) (*Repositories, error) {
	postgresDB := db.(*sqlx.DB)
	switch dbType {
	case "postgres":
		return &Repositories{
			UserRepository:     NewUserRepository(postgresDB),
			ProductRepository:  NewProductRepository(postgresDB),
			DataBaseRepository: NewDataBaseRepository(postgresDB),
		}, nil
	default:
		return nil, errors.New("unsupported db type")
	}
}
