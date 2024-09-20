package pkggomigrations

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
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
