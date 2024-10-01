# TODOs for Project: go-migrations

## Project Description:
This project is a CLI tool for managing database migrations for PostgreSQL, developed in Go. In addition to its command-line interface (CLI) usage, the project is also designed to be used as a library by other Go software.

## Current Project Structure:

```
.
├── cmd
│   └── go-migrations
│       └── main.go
├── configs.yaml
├── go.mod
├── go.sum
├── internal
│   └── db
│       └── connect.go
├── LICENSE
├── migrations
│   ├── 0001_create_users_table.down.sql
│   ├── 0001_create_users_table.up.sql
│   └── applied
├── pkg-go-migrations
│   ├── migrates.go
│   ├── read_files.go
│   ├── utilities.go
│   └── write_files.go
├── postgres.sh
├── project.toml
├── README.md
└── TODO.md

7 directories, 16 files

```

---

## Pending Tasks:

### Features
- [X] Implement support for the `migrate up` command in the CLI.
  - [X] Read and apply migrations from the `pkg-go-migrations/migrations` directory.
  - [X] Display success or failure messages after each migration.
- [X] Implement support for the `migrate down` command in the CLI.
  - [ ] Add rollback for the last applied migrations.
- [X] Add support for multiple databases (PostgreSQL, MySQL).
- [ ] Create an interface for integrating the package with other Go applications.
  - [ ] Export migration functions for use in other software.

### Structure and Organization
- [ ] Refactor the project structure:
  - [X] Create a `migrations` folder to store migration files.
  - [ ] Modularize the migration functions in the `pkg-go-migrations/` directory.
- [ ] Write unit tests for the migration functions.
  - [ ] Test `up` migrations.
  - [ ] Test `down` migration rollbacks.

### Documentation
- [ ] Add a "How to use" section to `README.md` to explain the CLI commands.
  - [ ] Explain how to use `migrate up` and `migrate down`.
  - [ ] Provide instructions for integration with other Go projects.
- [ ] Create detailed documentation for the exported methods in the `pkg-go-migrations` package.
- [X] Write functions documentations.

### Future Improvements
- [ ] Implement support for automatic migrations based on schema files.
- [ ] Add audit logs for each migration executed.
- [ ] Create a graphical interface to facilitate migrations via the web (future).

### Micro-tasks:
#### File Handling
- [X] Read YAML configuration files.
- [X] Read TOML configuration files.
- [X] Read SQL migration files.
- [X] Create path migrations.
- [X] Create SQL files for migrations (up/down).
- [ ] Create general migration files (up/down).

#### Database Operations
- [X] Create PostgreSQL connections.
- [X] Create MySQL connections.
- [ ] Test and validate database connections.

#### Migrations Management
- [ ] Implement error handling for migration failures.
- [X] Log success or failure of migrations.
- [ ] Implement rollback for migrations.
- [ ] Handle concurrent migrations execution.
  
#### CLI Features
- [ ] Usage Flags
  - [X] Version
  - [X] Help
  - [X] Check Configs - Alter to method CheckDbConfig in *_checkConfig
  - [X] migrate-up - Add Mkdir applied
  - [X] migrate-down
  - [ ] migrate-ls
  - [X] Ping DB
- [ ] Display detailed logs on the CLI for each step of the migration.
- [ ] Provide CLI options for dry-run migrations.

#### Tests
- [ ] read files
  - [X] ReadProjectToml
  - [ ] ReadYamlConfig 
  - [ ] ReadMigration 


---

## Priorities:

### High Priority:
1. Complete the implementation of the `migrate up` and `migrate down` commands.
2. Refactor the structure to support migrations for multiple databases.

### Low Priority:


---

This document should be updated as the development progresses.