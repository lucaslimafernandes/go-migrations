package pkggomigrations

func CheckDbConfigEmpty(db *DBConfig) map[string]bool {

	res := make(map[string]bool)

	if db.Postgres.Host != "" {
		res["postgres"] = true
	} else {
		res["postgres"] = false
	}

	if db.Mysql.Host != "" {
		res["mysql"] = true
	} else {
		res["mysql"] = false
	}

	return res

}
