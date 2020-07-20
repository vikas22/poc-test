// Application execution starts here
// This will have all blank imports used in the application as part of convention
// This will initialize the application based on the parameter specified
package main

import (
	"db-poc/internal/bootstrap"
	"flag"
	"fmt"
	"log"
	"os"

	_ "db-poc/internal/migration"
	_ "github.com/jinzhu/gorm/dialects/mysql" // We use mysql dialect as database and this blank import, refer database/sql.
	"github.com/pressly/goose"
)

const (
	MigrationDir = "dir"

	DefaultMigrationDir = "internal/migration"
)

var (
	flags = flag.NewFlagSet("migrate", flag.ExitOnError)
	dir   = flags.String(MigrationDir, DefaultMigrationDir, "directory with migration files")
)

const (
	// CREATE - CREATE COMMAND
	CREATE = "create"
)

// Execution starts here
// while execution we can give the command line parameters to configure the application
// There are two flags which can be set while initiating
// - base_path: default it'll take the current directory as application
//			    configuration will be loaded from `conf` folder available in the directory specified by this
// - env: default it'll be `dev`. application will be initialized according to this
// - command: when specified application will initialize in command mode and execute the given command
//			  router wont be initialized for commands
func main() {
	usage()
	flags.Parse(os.Args[1:])
	args := flags.Args()

	// I.e. no command provided, hence print usage and return.
	if len(args) < 1 {
		flags.Usage()
		return
	}

	command := args[0]
	arguments := args[1:]

	if len(args) > 1 && command == CREATE {
		if err := goose.Run(CREATE, nil, *dir, arguments...); err != nil {
			log.Fatalf("goose run: %v", err)
		}

		return
	}

	err := bootstrap.Init()

	if err != nil {
		panic(err)
	}

	dialect := bootstrap.Config.Database.Dialect
	if err := goose.SetDialect(dialect); err != nil {
		log.Fatalf("Migration: Failed to run command: %v", err)
	}
	sqldb := bootstrap.Db.Instance().DB()

	// Finally, executes the goose's command.
	if err := goose.Run(command, sqldb, *dir, args[1:]...); err != nil {
		log.Fatalf("Migration: Failed to run command: %v", err)
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
}

var (
	usagePrefix = `Usage: goose [OPTIONS] DRIVER DBSTRING COMMAND
Drivers:
    postgres
    mysql
    sqlite3
    redshift
Examples:
    goose sqlite3 ./foo.db status
    goose sqlite3 ./foo.db create init sql
    goose sqlite3 ./foo.db create add_some_column sql
    goose sqlite3 ./foo.db create fetch_user_data go
    goose sqlite3 ./foo.db up
    goose postgres "user=postgres dbname=postgres sslmode=disable" status
    goose mysql "user:password@/dbname?parseTime=true" status
    goose redshift "postgres://user:password@qwerty.us-east-1.redshift.amazonaws.com:5439/db" status
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
