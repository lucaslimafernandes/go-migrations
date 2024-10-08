package pkggomigrations

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

// Migrate applies a migration to the database, either up or down based on the 'upDown' parameter.
// It executes the SQL script corresponding to the migration version and direction (up/down).
func Migrate(version string, db *sql.DB, upDown string, pathMigrations string) {

	fi, flname, _ := ReadMigration(version, upDown, pathMigrations)

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

	err = write_in(&fi, flname, int(rowsAff), int(lastId), pathMigrations)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Migration applied sucessfully: ", flname)

}
