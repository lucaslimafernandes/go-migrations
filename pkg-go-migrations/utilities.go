package pkggomigrations

import "fmt"

type VerifyDbConfig struct {
	Migrations struct {
		PathMigrations bool `yaml:"PATH"`
	} `yaml:"migrations"`
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

// CheckDbConfigEmpty checks if the database configurations for Postgres and MySQL are set.
// It returns a map of booleans indicating whether the `Host` field is filled for each database.
//
// Params:
// - db: A pointer to the DBConfig structure that holds the configuration for Postgres and MySQL.
//
// Returns:
//   - map[string]bool: A map with two keys, "postgres" and "mysql", where each key indicates if
//     the respective database's Host field is filled (true) or empty (false).
//
// Example usage:
//
//	config := &DBConfig{...}
//	res := CheckDbConfigEmpty(config)
//	fmt.Println(res["postgres"]) // true if Postgres Host is set
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

// CheckDbConfigApply verifies if the `Apply` flag is set for Postgres and MySQL in the configuration.
// It returns a map indicating whether migrations should be applied to each database, and a boolean
// value indicating if the configuration is valid for one database at a time.
//
// Params:
// - db: A pointer to the DBConfig structure that holds the configuration for Postgres and MySQL.
//
// Returns:
//   - map[string]string: A map with two keys, "true" and "false", where each key indicates if
//     the respective database's Apply field is "postgres" or "mysql".
//   - bool: A boolean value that is true if exactly one of Postgres or MySQL has Apply set to true.
//
// Example usage:
//
//	config := &DBConfig{...}
//	res, valid := config.CheckDbConfigApply()
//	fmt.Println(res["true"], valid)
func (db *DBConfig) CheckDbConfigApply() (map[string]string, bool) {

	var valid bool
	res := make(map[string]string)

	if db.Postgres.Apply {
		res["true"] = "postgres"
	} else {
		res["false"] = "postgres"
	}

	if db.Mysql.Apply {
		res["true"] = "mysql"
	} else {
		res["false"] = "mysql"
	}

	if db.Postgres.Apply && !db.Mysql.Apply {
		valid = true
	} else if db.Mysql.Apply && !db.Postgres.Apply {
		valid = true
	}

	return res, valid

}

// CheckDbConfig validates the configuration for Postgres and MySQL by checking various fields such as
// `Host`, `Port`, `User`, and `Password`, and prints a formatted summary of the validation results.
//
// This function also checks if the configuration is valid by ensuring only one of Postgres or MySQL
// has the `Apply` flag set to true and the path for migrations.
//
// Params:
// - config: A pointer to the DBConfig structure that holds the configuration for Postgres and MySQL.
//
// Prints:
//   - A formatted output detailing whether each field (Apply, Host, Port, User, Password) is valid for
//     both Postgres and MySQL.
//   - Whether the configuration is valid, meaning only one database has the `Apply` flag set to true.
//
// Example output:
//
//	Check configuration DB connect file
//
// Path for migrations is ok: true
//
//	/path/migrations/
//
//	Postgres:
//	  APPLY: true
//	  HOST: true
//	  PORT: true
//	  USER: true
//	  PASSWORD: true
//	Mysql:
//	  APPLY: false
//	  HOST: false
//	  PORT: false
//	  USER: false
//	  PASSWORD: false
//	Is Valid: true
//	go-migrations accept only one DB at a time
//
// Example usage:
//
//	config := &DBConfig{...}
//	config.CheckDbConfig()
func (config *DBConfig) CheckDbConfig() {

	var res VerifyDbConfig

	_, isValid := config.CheckDbConfigApply()

	//
	// Verification Migrations Path
	//

	if config.Migrations.PathMigrations != "" {
		res.Migrations.PathMigrations = true
	} else {
		res.Migrations.PathMigrations = false
	}

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

	fmt.Printf(`Check configuration DB connect file

Path for migrations is ok: %v
	%v

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
		res.Migrations.PathMigrations, config.Migrations.PathMigrations,
		res.Postgres.Apply, res.Postgres.Host, res.Postgres.Port, res.Postgres.User, res.Postgres.Password,
		res.Mysql.Apply, res.Mysql.Host, res.Mysql.Port, res.Mysql.User, res.Mysql.Password, isValid,
	)

}
