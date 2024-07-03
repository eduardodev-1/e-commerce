package repositories

import (
	"github.com/jmoiron/sqlx"
)

type DataBaseRepository struct {
	*sqlx.DB
}

func NewDataBaseRepository(db *sqlx.DB) *DataBaseRepository {
	return &DataBaseRepository{
		db,
	}
}
