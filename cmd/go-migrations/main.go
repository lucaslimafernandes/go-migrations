package main

import (
	"flag"
	"fmt"
	"log"
	"os"

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

	// test
	// p, e := pkggomigrations.ReadYamlConfig("configs.yaml")
	// fmt.Println(p)
	// fmt.Println(e)

	// fmt.Println(pkggomigrations.CheckDbConfigEmpty(p))
	// fmt.Println(pkggomigrations.CheckDbConfigApply(p))
	// test 2
	// n := "0001"
	// ls, _ := pkggomigrations.ReadMigration(n, ".up.sql")
	// fmt.Println(ls)

}

func handler() {

	_version := flag.Bool("version", false, "Print version of the go-migrations")
	_help := flag.Bool("help", false, "Show available commands")
	_checkConfig := flag.Bool("check-config", false, "Verify the yaml file")

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
		p, e := pkggomigrations.ReadYamlConfig("configs.yaml")
		if e != nil {
			fmt.Printf("%v - version: %v\n", projectToml.Project.Name, projectToml.Project.Version)
			fmt.Println(e)
		} else {
			verif := pkggomigrations.CheckDbConfig(p)
			fmt.Printf("%v - version: %v\n", projectToml.Project.Name, projectToml.Project.Version)
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

`,
				verif.Postgres.Apply, verif.Postgres.Host, verif.Postgres.Port, verif.Postgres.User, verif.Postgres.Password,
				verif.Mysql.Apply, verif.Mysql.Host, verif.Mysql.Port, verif.Mysql.User, verif.Mysql.Password,
			)
		}

		return
	}

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
		return
	}

}
