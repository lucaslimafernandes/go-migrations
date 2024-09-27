package pkggomigrations

import "fmt"

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

// func CheckDbConfigApply(db *DBConfig) map[string]bool {
func (db *DBConfig) CheckDbConfigApply() (map[string]bool, bool) {

	var valid bool
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

	if db.Postgres.Apply && !db.Mysql.Apply {
		valid = true
	} else if db.Mysql.Apply && !db.Postgres.Apply {
		valid = true
	}

	return res, valid

}

func (config *DBConfig) CheckDbConfig() {

	var res VerifyDbConfig

	_, isValid := config.CheckDbConfigApply()

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

	// return res

	fmt.Printf(`Check configuration DB connect file

Postgres:
	APPLY: %v
	HOST: %v
	PORT: %v
	USER: %v
	PASSWORD: %v
Mysql:
	APPLY: %v
	HOST: %v
	PORT: %v
	USER: %v
	PASSWORD: %v

Is Valid: %v
go-migrations accept only one DB at a time 

`,
		res.Postgres.Apply, res.Postgres.Host, res.Postgres.Port, res.Postgres.User, res.Postgres.Password,
		res.Mysql.Apply, res.Mysql.Host, res.Mysql.Port, res.Mysql.User, res.Mysql.Password, isValid,
	)

}
