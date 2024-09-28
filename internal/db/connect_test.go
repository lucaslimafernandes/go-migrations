package db

import (
	"testing"

	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

// go test ./internal/db/
func TestDBConnect(t *testing.T) {

	dbConfig := pkggomigrations.DBConfig{
		Postgres: struct {
			Apply    bool   "yaml:\"APPLY\""
			Host     string "yaml:\"HOST\""
			Port     string "yaml:\"PORT\""
			User     string "yaml:\"USER\""
			Password string "yaml:\"PASSWORD\""
			Dbname   string "yaml:\"DBNAME\""
		}{
			Apply:    true,
			Host:     "localhost",
			Port:     "5432",
			User:     "postgres",
			Password: "password",
			Dbname:   "postgres",
		},
		Mysql: struct {
			Apply    bool   "yaml:\"APPLY\""
			Host     string "yaml:\"HOST\""
			Port     string "yaml:\"PORT\""
			User     string "yaml:\"USER\""
			Password string "yaml:\"PASSWORD\""
			Dbname   string "yaml:\"DBNAME\""
		}{
			Apply:    true,
			Host:     "localhost",
			Port:     "3306",
			User:     "myuser",
			Password: "user_password",
			Dbname:   "mydb",
		},
	}

	// Postgres
	name := "postgres"
	dbPg, err := DBConnect(dbConfig, name)
	if err != nil {
		t.Errorf("Failure %s: %v", name, err)
	}

	err = dbPg.Ping()
	if err != nil {
		t.Errorf("Failure ping %s: %v", name, err)
	}

	// MySql
	name = "mysql"
	dbMy, err := DBConnect(dbConfig, name)
	if err != nil {
		t.Errorf("Failure %s: %v", name, err)
	}

	err = dbMy.Ping()
	if err != nil {
		t.Errorf("Failure ping %s: %v", name, err)
	}

}
