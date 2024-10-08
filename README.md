# go-migrations
A database migration tool.

## Usage

### Help

```bash
go run cmd/go-migrations/main.go -help
```

```bash
Usage of go-migrations <0.1.0>
  -check-config
    	Verify the yaml file
  -help
    	Show available commands
  -migrate-down
    	Make migrations down
  -migrate-up
    	Make migrations up
  -ping
    	Check the DB connection
  -v string
    	Specify the version ID for the migration (Format: 0001)
  -version
    	Print version of the go-migrations

```

### Check-config

```bash
go run cmd/go-migrations/main.go -check-config
go-migrations - version: 0.1.0
Check configuration DB connect file

Path for migrations is ok: true
	/home/lucas/go/src/github.com/lucaslimafernandes/go-migrations/migrations

Postgres:
	APPLY: true
	HOST: true
	PORT: true
	USER: true
	PASSWORD: true
Mysql:
	APPLY: false
	HOST: true
	PORT: true
	USER: true
	PASSWORD: true

Is Valid: true
go-migrations accept only one DB at a time 
```


### Migrate-up

```bash
go run cmd/go-migrations/main.go -migrate-up -v 0001
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

--Applied at: 2024-10-06 14:44:37
--Rows affected: 0
--Last Id inserted: 0
Migration applied sucessfully:  0001_create_users_table.up.sql
```
### Migrate-down

```bash
go run cmd/go-migrations/main.go -migrate-down -v 0001              
DROP TABLE if exists users;

--Applied at: 2024-10-06 14:46:09
--Rows affected: 0
--Last Id inserted: 0
Migration applied sucessfully:  0001_create_users_table.down.sql
```

### Ping

```bash
go run cmd/go-migrations/main.go -ping
go-migrations - version: 0.1.0
Ping
Everything Ok!% 
```

### Version

```bash
go run cmd/go-migrations/main.go -version
go-migrations - version: 0.1.0
```

## Configuration file

The `configs.yaml` file is used to configure the database connection settings and migration paths for the go-migrations tool. It defines settings for both PostgreSQL and MySQL databases, allowing the user to configure which database to apply migrations to. Below is a detailed explanation of each section in the configs.yaml file:

1. Migrations

- PATH: Specifies the directory where the migration files are stored. The migration files define the database schema changes and are applied in sequence. Ensure that this path points to the correct directory where the .sql migration files (e.g., .up.sql and .down.sql) are located.

2. PostgreSQL Configuration

- APPLY: Set to true to apply migrations to the PostgreSQL database.
- HOST: The hostname or IP address of the PostgreSQL server. Defaults to localhost if running on the same machine.
- PORT: The port number to connect to PostgreSQL. The default PostgreSQL port is 5432.
- USER: The username to use when connecting to the PostgreSQL database.
- PASSWORD: The password for the specified PostgreSQL user.
- DBNAME: The name of the PostgreSQL database where migrations will be applied.

3. MySQL Configuration

- APPLY: Set to true to apply migrations to the MySQL database. Set to false if PostgreSQL is being used instead.
- HOST: The hostname or IP address of the MySQL server. Defaults to localhost.
- PORT: The port number to connect to MySQL. The default MySQL port is 3306.
- USER: The username to use when connecting to the MySQL database.
- PASSWORD: The password for the specified MySQL user.
- DBNAME: The name of the MySQL database where migrations will be applied.
- ROOT_PASSWORD: The root password for the MySQL server.

Note: To run migrations on PostgreSQL, ensure that APPLY is set to true. If you are only using MySQL, set APPLY to false for PostgreSQL.

To run migrations on PostgreSQL, ensure that:

  - postgres.APPLY is set to true.
  - mysql.APPLY is set to false.

To run migrations on MySQL, ensure that:

  - mysql.APPLY is set to true.
  - postgres.APPLY is set to false.

## Contributing

Your contributions are welcome! If you encounter any bugs or have feature requests, please open an issue. To contribute code, follow these steps:

1. Fork the repository.
2. Clone your forked repository to your local machine.
3. Create a new branch (git checkout -b feature-or-bugfix-name).
4. Make your changes and commit them (git commit -m "Description of your changes").
5. Push your branch to your forked repository (git push origin feature-or-bugfix-name).
6. Open a pull request with a clear description of your changes.
