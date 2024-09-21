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

	_version := flag.Bool("version", false, "Print version of the go-migrations")
	_help := flag.Bool("help", false, "Show available commands")

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

	// test
	p, e := pkggomigrations.ReadYamlConfig("configs.yaml")
	fmt.Println(p)
	fmt.Println(e)

	fmt.Println(pkggomigrations.CheckDbConfigEmpty(p))
	fmt.Println(pkggomigrations.CheckDbConfigApply(p))
	// test 2
	n := "0001"
	ls, _ := pkggomigrations.ReadMigration(n, ".up.sql")
	fmt.Println(ls)

	if len(os.Args) == 1 {
		fmt.Println("Maybe you forget some commands, use '-help' to see available commands.")
	}

}
