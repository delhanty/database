package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	_ "git-codecommit.us-east-1.amazonaws.com/v1/repos/database/migrations"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

const (
	driver = "postgres"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	args := flags.Args()

	if len(args) > 1 && args[0] == "create" {
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	log.Printf("%#v\n", args)
	if args[0] == "-h" || args[0] == "--help" {
		flags.Usage()
		return
	}

	command := args[0]

	if err := goose.SetDialect(driver); err != nil {
		log.Fatalf("unable to set postgres dialect: %s\n", err)
	}

	dbstring := os.Getenv("GOOSE_DB_STRING")
	if dbstring == "" {
		log.Fatalf("-dbstring=%q not supported\n", dbstring)
	}

	db, err := sql.Open(driver, dbstring)
	if err != nil {
		log.Fatalf("-dbstring=%q: %v\n", dbstring, err)
	}

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose run: %v", err)
	}
}

func usage() {
	log.Print(usagePrefix)
	flags.PrintDefaults()
	log.Print(usageCommands)
}

var (
	usagePrefix = `Usage: goose [OPTIONS] COMMAND
Examples:
	(export GOOSE_DB_STRING="user=postgres dbname=postgres sslmode=disable")
	goose -dir ./migrations status
	goose -dir ./migrations create init sql
	goose -dir ./migrations create something_from_go_file go
	goose -dir ./migrations up
	goose -dir ./migrations down
	goose -dir ./migrations redo
Options:
`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with next version
`
)
