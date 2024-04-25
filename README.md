# Example of DB Migration Mechanism Using [sqldef](https://github.com/sqldef/sqldef)

You can use the command presets prepared in `Taskfile.yml` to easily manage the DB schema. The current schema is written by SQL in `tool/postgres/schema.sql`. If you want to change the running DB's schema, simply edit it directly and run the specified apply command.

## Prerequisites
Before you begin, ensure you have met the following requirements:

- [direnv](https://github.com/direnv/direnv): Used to load and unload environment variables depending on the current directory. This is crucial for managing project-specific configurations seamlessly.
- [Task](https://taskfile.dev/): A simple and flexible task runner written in Go, akin to a Makefile but with no dependencies and cross-platform support. It's used to automate common development tasks.
- [Docker](https://www.docker.com/): Required for containerization of the application, ensuring consistent environments across development, testing, and production.
Please follow the installation instructions linked above to set up these tools on your machine.

## Usage
### Command Presets in Taskfile.yml

#### Setup containers for PostgreSQL and sqldef:
These commands initialize and build the necessary Docker containers for PostgreSQL and sqldef, setting up your development environment for database operations.
```
$ task db:init
$ task sqldef:build
```

#### Export current schema:
This command outputs the current database schema to standard output. It allows you to view the current schema directly from your command line, which can be useful for quick reviews or documentation purposes.
```
$ task db:dump
```

#### Dry-run apply schema:
This command performs a dry-run of the schema application process, displaying the SQL statements that would be executed if the changes were applied. It does not make any actual changes to the database, serving as a safe way to verify the intended modifications.
```
$ task db:dryrun
```

#### Apply schema:
This command actually applies the schema changes to your database by executing the SQL statements. Use it to implement your schema changes after confirming them with a dry-run.
```
$ task db:apply
```
