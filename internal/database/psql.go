package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect1(databaseURL string) (*pgxpool.Pool, error) {
	var ctx context.Context = context.Background()

	var config *pgxpool.Config
	var err error

	if config, err = pgxpool.ParseConfig(databaseURL); err != nil {
		log.Printf("Unable to parse DATABSE_URL: %v", err)
		return nil, err
	}

	var pool *pgxpool.Pool

	if pool, err = pgxpool.NewWithConfig(ctx, config); err != nil {
		log.Printf("Unable to create connection pool: %v", err)
		pool.Close()
		return nil, err
	}

	if err = pool.Ping(ctx); err != nil {
		log.Printf("Unable to ping database: %v", err)
		return nil, err
	}

	log.Println("Successfully connected to postgreSQL databse")
	return pool, nil

}
