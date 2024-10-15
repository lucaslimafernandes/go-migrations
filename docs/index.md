# go-migrations
A database migration tool.


Go-migrations is a CLI tool and Go library designed to simplify and manage database migrations for projects that use PostgreSQL or MySQL. It provides an easy-to-use interface for applying, reverting, and managing database schema changes across different environments, ensuring consistency and version control of database states.

This project is intended to be flexible, offering the following key features:

- Multi-Database Support: Supports both PostgreSQL and MySQL databases, allowing seamless transitions between different database systems as needed.

- CLI for Migration Management: Provides a command-line interface to apply migrations (up) or revert migrations (down), with version control to ensure that migrations are applied in a consistent and correct order.

- Configurable via configs.yaml: The project uses a configuration file to specify database connection details and the location of migration files, ensuring flexibility across different environments (e.g., development, staging, production).

- Versioned Migrations: Each migration file follows a naming convention that includes version numbers, ensuring that changes to the database schema are tracked and can be rolled back or reapplied as necessary.

The primary goal of go-migrations is to provide developers with an efficient and reliable way to manage database migrations, preventing issues related to inconsistent database states and ensuring that changes can be easily rolled back if necessary. Whether working in a microservice architecture, monolithic application, or across multiple database environments, go-migrations provides the tools needed to keep database schemas under control.

## Usage

[How to use CLI Go-migrations](usage.md)


## Configuration file

[How to use CLI Go-migrations](configuration_file.md)
