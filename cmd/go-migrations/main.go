package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/lucaslimafernandes/go-migrations/internal/db"
	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

var projectToml *pkggomigrations.ProjectToml

func init() {

	var err error
	projectToml, err = pkggomigrations.ReadProjectToml()

	if err != nil {
		log.Fatalln("Failed to initialize go-migrations!")
	}

}

func main() {

	handler()

}

// handler is the main entry point for command-line arguments parsing in the go-migrations tool.
// It handles different flags passed via the command line and triggers corresponding actions such as
// printing the version, showing available commands, checking the YAML configuration, and managing database migrations.
//
// Flags:
//
//	-version: Prints the current version of the go-migrations tool.
//	-help: Displays the available commands and usage information for the tool.
//	-check-config: Verifies the validity of the configuration in the 'configs.yaml' file.
//	-migrate-up: Runs migrations to move the database schema upwards to a specific version.
//	-migrate-down: Reverts migrations by moving the database schema downwards to a previous version.
//	-v: Specifies the version ID for the migration to be applied (e.g., "0001").
//
// The function follows this logic:
//  1. If '-version' is passed, it prints the name and version of the tool.
//  2. If '-help' is passed, it prints the usage of the tool and the available flags.
//  3. If '-check-config' is passed, it verifies and validates the YAML configuration file.
//  4. If either '-migrate-up' or '-migrate-down' is passed along with a version ID ('-v'),
//     it performs the corresponding migration using the version ID specified.
//
// Errors:
//   - If both '-migrate-up' and '-migrate-down' are used together, the tool will print an error message.
//   - If no version is provided with '-migrate-up' or '-migrate-down', the tool prints a message suggesting the use of help.
//   - If no valid arguments are passed, the tool prompts the user to use '-help'.
func handler() {

	_version := flag.Bool("version", false, "Print version of the go-migrations")
	_help := flag.Bool("help", false, "Show available commands")
	_checkConfig := flag.Bool("check-config", false, "Verify the yaml file")
	_migrateUp := flag.Bool("migrate-up", false, "Make migrations up")
	_migrateDown := flag.Bool("migrate-down", false, "Make migrations down")
	_migrateVersion := flag.String("v", "", "Specify the version ID for the migration (Format: 0001)")

	flag.Parse()

	if *_version {
		fmt.Printf("%v - version: %v\n", projectToml.Project.Name, projectToml.Project.Version)
		return
	}

	if *_help {
		fmt.Printf("Usage of %v <%v>", projectToml.Project.Name, projectToml.Project.Version)
		flag.PrintDefaults()
		return
	}

	if *_checkConfig {

		var p *pkggomigrations.DBConfig

		p, e := pkggomigrations.ReadYamlConfig("configs.yaml")
		if e != nil {
			fmt.Printf("%v - version: %v\n", projectToml.Project.Name, projectToml.Project.Version)
			fmt.Println(e)
		} else {
			fmt.Printf("%v - version: %v\n", projectToml.Project.Name, projectToml.Project.Version)
			p.CheckDbConfig()

			res, isValid := p.CheckDbConfigApply()
			fmt.Printf("\n%v\n%v\n", isValid, res)

		}

		return
	}

	if *_migrateUp || *_migrateDown {

		if *_migrateVersion == "" {
			fmt.Println("Maybe you need use help '-help'.")
			return
		}

		if *_migrateUp && *_migrateDown {
			fmt.Println("One type migration at a time")
		}

		p, _ := pkggomigrations.ReadYamlConfig("configs.yaml")

		_, isValid := p.CheckDbConfigApply()
		if !isValid {
			fmt.Println("Maybe you need check the config file, use '-check-config'.")
			return
		}

		database, err := db.PgConnect(*p)
		if err != nil {
			log.Fatalln(err)
		}

		if *_migrateUp && !*_migrateDown {
			pkggomigrations.Migrate(*_migrateVersion, database, ".up.sql")
		}

		if !*_migrateUp && *_migrateDown {
			pkggomigrations.Migrate(*_migrateVersion, database, ".down.sql")
		}

	}

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
		return
	}

}
