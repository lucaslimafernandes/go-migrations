package pkggomigrations

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
)

type ProjectToml struct {
	Project struct {
		Name        string
		Version     string
		Description string
	}
	Developer struct {
		Author string
		Email  string
	}
	Repository struct {
		URL string
	}
}

type PostgresConfig struct {
	Postgres struct {
		Host     string `yaml:"HOST"`
		Port     string `yaml:"PORT"`
		User     string `yaml:"USER"`
		Password string `yaml:"PASSWORD"`
	} `yaml:"postgres"`
}

func ReadProjectToml() (*ProjectToml, error) {

	var pfile ProjectToml

	f, err := os.ReadFile("project.toml")
	if err != nil {
		log.Fatalf("Failed to read project docs: %v", err)
	}

	_, err = toml.Decode(string(f), &pfile)
	if err != nil {
		log.Fatalf("Failed to Unmarshal TOML: %v", err)
		return nil, err
	}

	return &pfile, nil

}

func ReadYamlConfig(filename string) (*PostgresConfig, error) {

	var pg PostgresConfig
	var err error

	f, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read configs: %v", err)
		return nil, err
	}

	err = yaml.Unmarshal(f, &pg)
	if err != nil {
		log.Printf("Some configuration need attention!")
		return nil, err
	}

	return &pg, nil

}
