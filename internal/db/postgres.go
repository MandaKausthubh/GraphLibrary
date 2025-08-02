package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)


var Pool *pgxpool.Pool

func InitPostgres() {
	url := os.Getenv("POSTGRES_URL")
	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		panic("Failed to parse Postgres URL: " + err.Error())
	}

	config.MaxConns = 10
	config.MinConns = 2
	config.MaxConnIdleTime = time.Hour

	Pool, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		panic("Failed to connect to Postgres: " + err.Error())
	}
	if err := Pool.Ping(context.Background()); err != nil {
		panic("Failed to ping Postgres: " + err.Error())
	}
	log.Println("🆗: Connected to Postgres successfully")
}
