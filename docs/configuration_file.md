# Configuration

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
