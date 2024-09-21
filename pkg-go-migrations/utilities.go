package pkggomigrations

type VerifyDbConfig struct {
	Postgres struct {
		Apply    bool `yaml:"APPLY"`
		Host     bool `yaml:"HOST"`
		Port     bool `yaml:"PORT"`
		User     bool `yaml:"USER"`
		Password bool `yaml:"PASSWORD"`
	} `yaml:"postgres"`
	Mysql struct {
		Apply    bool `yaml:"APPLY"`
		Host     bool `yaml:"HOST"`
		Port     bool `yaml:"PORT"`
		User     bool `yaml:"USER"`
		Password bool `yaml:"PASSWORD"`
	} `yaml:"mysql"`
}

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

func CheckDbConfigApply(db *DBConfig) map[string]bool {

	res := make(map[string]bool)

	if db.Postgres.Apply {
		res["postgres"] = true
	} else {
		res["postgres"] = false
	}

	if db.Mysql.Apply {
		res["mysql"] = true
	} else {
		res["mysql"] = false
	}

	return res

}

func CheckDbConfig(config *DBConfig) VerifyDbConfig {

	var res VerifyDbConfig

	//
	// Verification Postgres
	//

	if config.Postgres.Apply {
		res.Postgres.Apply = true
	} else {
		res.Postgres.Apply = false
	}

	if config.Postgres.Host != "" {
		res.Postgres.Host = true
	} else {
		res.Postgres.Host = false
	}

	if config.Postgres.Password != "" {
		res.Postgres.Password = true
	} else {
		res.Postgres.Password = false
	}

	if config.Postgres.Port != "" {
		res.Postgres.Port = true
	} else {
		res.Postgres.Port = false
	}

	if config.Postgres.User != "" {
		res.Postgres.User = true
	} else {
		res.Postgres.User = false
	}

	//
	// Verification Mysql
	//

	if config.Mysql.Apply {
		res.Mysql.Apply = true
	} else {
		res.Mysql.Apply = false
	}

	if config.Mysql.Host != "" {
		res.Mysql.Host = true
	} else {
		res.Mysql.Host = false
	}

	if config.Mysql.Password != "" {
		res.Mysql.Password = true
	} else {
		res.Mysql.Password = false
	}

	if config.Mysql.Port != "" {
		res.Mysql.Port = true
	} else {
		res.Mysql.Port = false
	}

	if config.Mysql.User != "" {
		res.Mysql.User = true
	} else {
		res.Mysql.User = false
	}

	return res

}
