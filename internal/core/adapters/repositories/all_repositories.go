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
	UserRepository     ports.UserRepository
	ProductRepository  ports.ProductRepository
	CategoryRepository ports.CategoryRepository
}

func NewRepositories(db *models.DataBase) *Repositories {
	switch db.Type {
	case Postgresql:
		postgresDB := db.Db.(*sqlx.DB)
		return NewPostgreeRepositories(postgresDB)
	case Mongodb:
		return NewMongoDBRepositories(db)
	default:
		log.Fatal("unsupported db type")
	}
	return nil
}
func NewPostgreeRepositories(postgresDB *sqlx.DB) *Repositories {
	return &Repositories{
		UserRepository:     postgres.NewUserRepository(postgresDB),
		ProductRepository:  postgres.NewProductRepository(postgresDB),
		CategoryRepository: postgres.NewCategoryRepository(postgresDB),
	}
}
func NewMongoDBRepositories(db *models.DataBase) *Repositories {
	return nil
}
