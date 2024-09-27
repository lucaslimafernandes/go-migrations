package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

// PgConnect establishes a connection to a PostgreSQL database using the provided configuration.
// It returns a pointer to the sql.DB object representing the connection and an error if something goes wrong.
func PgConnect(pgConfig pkggomigrations.DBConfig) (*sql.DB, error) {

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=postgres sslmode=disable", pgConfig.Postgres.Host, pgConfig.Postgres.User, pgConfig.Postgres.Password)

	// enhancement: dbdriver to DBConfig
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println("Failed to connect postgres: ", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Failed to ping to postgres: ", err)
		return nil, err
	}

	return db, nil

}
