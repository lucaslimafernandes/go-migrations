package main

import (
	"flag"
	"fmt"
)

func main() {

	_version := flag.Bool("version", false, "Print version of th go-migrations")

	flag.Parse()

	if *_version {
		fmt.Println("go-migrations - version: v0.0.0")
	}

}
