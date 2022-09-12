package app

import (
	"context"

	"github.com/JustinJohnsonK/will-share/pkg/log"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func intiDb() *pgxpool.Pool {
	poolConfig, err := pgxpool.ParseConfig(Config.Db.ConnectionString())
	if err != nil {
		Logger.Fatal("Unable to parse DBPostgresAddr", err, log.Fields{"address": Config.Db.Host})
	}

	poolConfig.ConnConfig.LogLevel = pgx.LogLevelWarn

	db, err := pgxpool.ConnectConfig(context.Background(), poolConfig)
	if err != nil {
		Logger.Fatal("Unable to create connection pool", err, log.Fields{"db_config": Config.Db})
	}

	return db
}
