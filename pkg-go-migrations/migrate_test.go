package pkggomigrations_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/lucaslimafernandes/go-migrations/internal/db"
	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

func TestMigrate(t *testing.T) {

	_migrateVersion := "T0001"

	p, _ := pkggomigrations.ReadYamlConfig("../configs.yaml")

	wDb, isValid := p.CheckDbConfigApply()
	if !isValid {
		fmt.Println("Maybe you need check the config file, use '-check-config'.")
		return
	}

	database, err := db.DBConnect(*p, wDb["true"])
	if err != nil {
		log.Fatalln(err)
	}

	pkggomigrations.Migrate(_migrateVersion, database, ".up.sql", p.Migrations.PathMigrations)

}
