package postgres

import (
	"e-commerce/internal/core/adapters/repositories"
	"e-commerce/internal/core/domain/models"
	"fmt"
	"log"
	"os"
	"strings"

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

func NewPsqlConn() *models.DB {
	params := dbParams{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PW"),
		DBName:   os.Getenv("DB_NAME"),
		SslMode:  os.Getenv("DB_SSL_MODE"),
	}

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		params.Host, params.Port, params.User, params.Password, params.DBName, params.SslMode)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	if err = executeSQLFile("internal/database/postgres/sql/schema.sql", db); err != nil {
		if containsAll(err.Error(), []string{"constraint", "for relation", "already exists"}) {
			log.Println("Constraint already exists, skipping...")
		} else {
			log.Fatal("Failed to execute schema.sql:", err)
		}
	}
	return &models.DB{
		Db:   db,
		Type: repositories.Postgresql,
	}
}
func executeSQLFile(filepath string, db *sqlx.DB) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	script := string(file)

	tx, err := db.Beginx()
	if err != nil {
		return err
	}

	result, err := tx.Exec(script)
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return errRollback
		}
		return err
	}
	fmt.Println(result)
	return tx.Commit()
}
func containsAll(mainStr string, substrs []string) bool {
	for _, substr := range substrs {
		if !strings.Contains(mainStr, substr) {
			return false
		}
	}
	return true
}
