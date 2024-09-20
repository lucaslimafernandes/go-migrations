package main

import (
	"flag"
	"fmt"
	"log"

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

	_version := flag.Bool("version", false, "Print version of th go-migrations")

	flag.Parse()

	if *_version {
		fmt.Printf("go-migrations - version: %v\n", projectToml.Project.Version)
	}

}
