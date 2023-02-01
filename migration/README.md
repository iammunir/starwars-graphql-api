# How to use Golang Migration

## Install Dependencies

Make sure your Go Environment has been setup correctly then running the following command

`go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

Run command below to check if Golang Migrate has been successfully installed

`migrate -help`

## Create Migration

```migrate create -ext sql -dir migration_file_name```

 - -ext sql, produces a sql file
 - -dir = define folder where to save the sql file 

CLI command above would produced 2 sql files: up and down, please write SQL command to create the table for UP and a command to drop the table for DOWN

## Execute a migration(s)

Make sure to create Database/Schema [if not exist] before executing the migration

`migrate -database "postgres://postgres:admin@localhost:5432/starwars_db?sslmode=disable&query" -path migration up`

