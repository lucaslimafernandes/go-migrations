package pkggomigrations

func MigrateUp(version string) {

	fi, _ := ReadMigration(version, ".up.sql")

	_ = write_in(&fi)

}
