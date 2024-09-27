package pkggomigrations

import (
	"fmt"
	"log"
	"os"
	"strings"

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

type DBConfig struct {
	Postgres struct {
		Apply    bool   `yaml:"APPLY"`
		Host     string `yaml:"HOST"`
		Port     string `yaml:"PORT"`
		User     string `yaml:"USER"`
		Password string `yaml:"PASSWORD"`
	} `yaml:"postgres"`
	Mysql struct {
		Apply    bool   `yaml:"APPLY"`
		Host     string `yaml:"HOST"`
		Port     string `yaml:"PORT"`
		User     string `yaml:"USER"`
		Password string `yaml:"PASSWORD"`
	} `yaml:"mysql"`
}

// ReadProjectToml read the file 'project.toml'
// It returns a pointer to the ProjectToml and an error if something goes wrong.
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

// ReadYamlConfig read the configuration file to connect in a database
// It returns a pointer to DBConfig and an error if something goes wrong.
func ReadYamlConfig(filename string) (*DBConfig, error) {

	var pg DBConfig
	var err error

	f, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Failed to read configs: %v\n", err)
		return nil, err
	}

	err = yaml.Unmarshal(f, &pg)
	if err != nil {
		log.Printf("Some configuration need attention!\n")
		return nil, err
	}

	return &pg, nil

}

// ReadMigration read the migration file to executes migrations based on the 'version' and 'mode' parameters.
// It returns the file content, the filename and an error if something goes wrong.
func ReadMigration(version string, mode string) (string, string, error) {

	var err error
	var migration string

	ls, err := os.ReadDir("migrations")
	if err != nil {
		log.Printf("Failed to load path 'migrations': %v\n", err)
		return "nil", "nil", err
	}

	for _, value := range ls {
		if strings.HasPrefix(value.Name(), version) && strings.HasSuffix(value.Name(), mode) {
			migration = value.Name()
			break
		}
	}

	f, err := os.ReadFile(fmt.Sprintf("migrations/%s", migration))
	if err != nil {
		log.Printf("Failed to read file (%s): %v\n", migration, err)
		return "nil", "nil", err
	}

	return string(f), migration, err
}
