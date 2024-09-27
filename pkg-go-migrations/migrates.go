package pkggomigrations

import (
	"context"
	"database/sql"
	"log"
)

func MigrateUp(version string, db *sql.DB) {

	fi, flname, _ := ReadMigration(version, ".up.sql")

	driver, err := db.Conn(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	result, err := driver.ExecContext(context.Background(), fi)
	if err != nil {
		log.Fatalln(err)
	}
	defer driver.Close()

	rowsAff, _ := result.RowsAffected()
	lastId, _ := result.LastInsertId()

	_ = write_in(&fi, flname, int(rowsAff), int(lastId))

}
