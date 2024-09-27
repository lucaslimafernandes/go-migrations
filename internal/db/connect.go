package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

// PgConnect establishes a connection to a database using the provided configuration.
// It returns a pointer to the sql.DB object representing the connection and an error if something goes wrong.
func DBConnect(pgConfig pkggomigrations.DBConfig, wDb string) (*sql.DB, error) {

	var connStr string

	if wDb == "postgres" {
		connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", pgConfig.Postgres.Host, pgConfig.Postgres.User, pgConfig.Postgres.Password, pgConfig.Postgres.Dbname)
	} else if wDb == "mysql" {
		connStr = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", pgConfig.Mysql.Host, pgConfig.Mysql.User, pgConfig.Mysql.Password, pgConfig.Mysql.Dbname)
	}

	db, err := sql.Open(wDb, connStr)
	if err != nil {
		log.Printf("Failed to connect %v: %v\n", wDb, err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Printf("Failed to ping to %v: %v\n", wDb, err)
		return nil, err
	}

	return db, nil

}
