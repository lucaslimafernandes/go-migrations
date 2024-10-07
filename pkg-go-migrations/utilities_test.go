package pkggomigrations_test

import (
	"reflect"
	"testing"

	pkggomigrations "github.com/lucaslimafernandes/go-migrations/pkg-go-migrations"
)

// TestCheckDbConfigEmpty tests different configurations of the DBConfig struct
// to verify if the CheckDbConfigEmpty function works correctly.
func TestCheckDbConfigEmpty(t *testing.T) {
	// Define test cases
	tests := []struct {
		name       string
		dbConfig   pkggomigrations.DBConfig
		wantResult map[string]bool
	}{
		{
			name: "Both Postgres and MySQL host filled",
			dbConfig: pkggomigrations.DBConfig{
				Postgres: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{
					Apply:    true,
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "password",
					Dbname:   "postgres",
				},
				Mysql: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{
					Apply:    true,
					Host:     "localhost",
					Port:     "3306",
					User:     "myuser",
					Password: "user_password",
					Dbname:   "mydb",
				},
			},
			wantResult: map[string]bool{
				"postgres": true,
				"mysql":    true,
			},
		},
		{
			name: "Only Postgres host filled",
			dbConfig: pkggomigrations.DBConfig{
				Postgres: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{
					Apply:    true,
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "password",
					Dbname:   "postgres",
				},
				Mysql: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{},
			},
			wantResult: map[string]bool{
				"postgres": true,
				"mysql":    false,
			},
		},
		{
			name: "No hosts filled",
			dbConfig: pkggomigrations.DBConfig{
				Postgres: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{},
				Mysql: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{},
			},
			wantResult: map[string]bool{
				"postgres": false,
				"mysql":    false,
			},
		},
	}

	// Run through each test case
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pkggomigrations.CheckDbConfigEmpty(&tt.dbConfig)
			if !reflect.DeepEqual(got, tt.wantResult) {
				t.Errorf("CheckDbConfigEmpty() = %v, want %v", got, tt.wantResult)
			}
		})
	}
}

func TestCheckDbConfigApply(t *testing.T) {

	tests := []struct {
		name       string
		dbConfig   pkggomigrations.DBConfig
		wantResult bool
	}{
		{
			name: "Both Postgres and MySQL apply is true - valid is false",
			dbConfig: pkggomigrations.DBConfig{
				Postgres: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{
					Apply:    true,
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "password",
					Dbname:   "postgres",
				},
				Mysql: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{
					Apply:    true,
					Host:     "localhost",
					Port:     "3306",
					User:     "myuser",
					Password: "user_password",
					Dbname:   "mydb",
				},
			},
			wantResult: false,
		},
		{
			name: "Only Postgres apply is true - valid is true",
			dbConfig: pkggomigrations.DBConfig{
				Postgres: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{
					Apply:    true,
					Host:     "localhost",
					Port:     "5432",
					User:     "postgres",
					Password: "password",
					Dbname:   "postgres",
				},
				Mysql: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{},
			},
			wantResult: true,
		},
		{
			name: "No applies filled",
			dbConfig: pkggomigrations.DBConfig{
				Postgres: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{},
				Mysql: struct {
					Apply    bool   "yaml:\"APPLY\""
					Host     string "yaml:\"HOST\""
					Port     string "yaml:\"PORT\""
					User     string "yaml:\"USER\""
					Password string "yaml:\"PASSWORD\""
					Dbname   string "yaml:\"DBNAME\""
				}{},
			},
			wantResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := tt.dbConfig.CheckDbConfigApply()
			if !reflect.DeepEqual(got, tt.wantResult) {
				t.Errorf("CheckDbConfigEmpty() = %v, want %v", got, tt.wantResult)
			}
		})
	}

}
