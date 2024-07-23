package repositories

import (
	"e-commerce/internal/core/adapters/repositories/postgres"
	"e-commerce/internal/core/domain/models"
	"e-commerce/internal/core/ports"
	"github.com/jmoiron/sqlx"
	"log"
)

const (
	Postgresql = "postgresql"
	Mongodb    = "mongodb"
)

type Repositories struct {
	UserRepository    ports.UserRepository
	ProductRepository ports.ProductRepository
}

func NewRepositories(db *models.DB) *Repositories {
	switch db.Type {
	case Postgresql:
		postgresDB := db.Db.(*sqlx.DB)
		return &Repositories{
			UserRepository:    postgres.NewUserRepository(postgresDB),
			ProductRepository: postgres.NewProductRepository(postgresDB),
		}
	case Mongodb:
		//instanciar repositories do mongodb, por exemplo
	default:
		log.Fatal("unsupported db type")
	}
	return nil
}
