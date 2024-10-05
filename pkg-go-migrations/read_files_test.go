package pkggomigrations_test

import (
	"testing"

	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

// TestReadProjectToml_Success tests the successful case of reading and parsing a TOML file.

func TestReadProjectToml(t *testing.T) {

	project, err := pkggomigrations.ReadProjectToml()
	if err != nil {
		t.Fatalf("failed to read project.toml: %v\n", err)
	}

	if project.Project.Name != "go-migrations" {
		t.Errorf("Expected project name 'go-migrations', got '%v'", project.Project.Name)
	}
	if project.Project.Version != "0.1.0" {
		t.Errorf("Expected version '0.1.0', got '%v'", project.Project.Version)
	}
	if project.Developer.Author != "Lucas Lima Fernandes" {
		t.Errorf("Expected author 'Lucas Lima Fernandes', got '%v'", project.Developer.Author)
	}
	if project.Repository.URL != "https://github.com/lucaslimafernandes/go-migrations" {
		t.Errorf("Expected URL 'https://github.com/lucaslimafernandes/go-migrations', got '%v'", project.Repository.URL)
	}

}

func TestReadYamlConfig(t *testing.T) {

}
