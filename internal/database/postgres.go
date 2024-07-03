package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dbParams struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func NewPsqlConn() (*sqlx.DB, error) {
	params := dbParams{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PW"),
		DBName:   os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", params.Host, params.Port, params.User, params.Password, params.DBName, params.SslMode)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
