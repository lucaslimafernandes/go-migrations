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

	// Test Conn PG
	// p, _ := pkggomigrations.ReadYamlConfig("configs.yaml")

	// _, isValid := p.CheckDbConfigApply()
	// if !isValid {
	// 	fmt.Println("Maybe you need check the config file, use '-check-config'.")
	// 	return
	// }

	// pkggomigrations.Write_in()

	// database, err := db.PgConnect(*p)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// driver, err := database.Conn(context.Background())
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// rows, err := driver.QueryContext(context.Background(), "SELECT 1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	var res int

	// 	_ = rows.Scan(&res)

	// 	fmt.Println("Row result: ", res)
	// }

}

func handler() {

	_version := flag.Bool("version", false, "Print version of the go-migrations")
	_help := flag.Bool("help", false, "Show available commands")
	_checkConfig := flag.Bool("check-config", false, "Verify the yaml file")
	_migrateUp := flag.Bool("migrate-up", false, "Make migrations up")

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

		}

		return
	}

	if *_migrateUp {

		p, _ := pkggomigrations.ReadYamlConfig("configs.yaml")

		_, isValid := p.CheckDbConfigApply()
		if !isValid {
			fmt.Println("Maybe you need check the config file, use '-check-config'.")
			return
		}

		pkggomigrations.Write_in()

	}

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
		return
	}

}
