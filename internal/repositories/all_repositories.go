package repositories

import (
	"e-commerce/internal/core/domain"
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

func NewRepositories(db *domain.DB) *Repositories {
	switch db.Type {
	case Postgresql:
		postgresDB := db.Db.(*sqlx.DB)
		return &Repositories{
			UserRepository:    NewUserRepository(postgresDB),
			ProductRepository: NewProductRepository(postgresDB),
		}
	case Mongodb:
		//instanciar repositories do mongodb, por exemplo
	default:
		log.Fatal("unsupported db type")
	}
	return nil
}
