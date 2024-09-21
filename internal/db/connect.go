package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

func PgConnect(pgConfig pkggomigrations.DBConfig) (*sql.DB, error) {

	connStr := fmt.Sprintf("user=%s password=%s dbname=postgres sslmode=disable", pgConfig.Postgres.User, pgConfig.Postgres.Password)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connect postgres: ", err)
		return nil, err
	}

	return db, nil

}
