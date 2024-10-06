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
go run cmd/go-migrations/main.go -migrate-down -v 0001              1 ↵
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
