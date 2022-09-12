package main

import (
	"database/sql"
	"flag"
	"os"

	"github.com/JustinJohnsonK/will-share/app"
	"github.com/JustinJohnsonK/will-share/pkg/log"
	_ "github.com/jackc/pgx/v4/stdlib"
	migrate "github.com/rubenv/sql-migrate"
)

func parseFlags() (string, int) {
	num := flag.Int("num", 0, "number of migrations to be rolled back")
	migrationType := flag.String("type", "up", "migrate 'up' or 'down'")
	flag.Parse()

	if *migrationType != "up" && *migrationType != "down" {
		flag.Usage()

		os.Exit(1)
	}

	return *migrationType, *num
}

func main() {
	env := os.Getenv("WILLSHARE-ENV")

	if env == "" {
		env = "dev"
	}

	app.LoadConfig(env)
	app.SetupLogger("willshare", "0.1")

	migrationType, num := parseFlags()

	connectionString := app.Config.Db.ConnectionString()
	db, err := sql.Open("pgx", connectionString)

	if err != nil {
		app.Logger.Fatal("unable to connect to database", err, log.Fields{})
	}
	defer db.Close()

	migrations := &migrate.FileMigrationSource{
		Dir: "internal/migrations",
	}

	switch migrationType {
	case "up":
		_, err = migrate.ExecMax(db, "postgres", migrations, migrate.Up, num)
	case "down":
		_, err = migrate.ExecMax(db, "postgres", migrations, migrate.Down, num)
	}

	if err != nil {
		app.Logger.Fatal("Error while running migrations", err, log.Fields{})
	}
}
