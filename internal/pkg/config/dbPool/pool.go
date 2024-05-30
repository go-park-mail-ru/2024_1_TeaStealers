package dbPool

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

var dbPool *pgxpool.Pool

func InitDatabasePool(databaseURL string, maxConns int32) {
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v\n", err)
	}
	config.MaxConns = maxConns
	config.MaxConnLifetime = time.Hour

	dbPool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
}

func GetDBPool() *pgxpool.Pool {
	return dbPool
}

func CloseDBPool() {
	dbPool.Close()
}
