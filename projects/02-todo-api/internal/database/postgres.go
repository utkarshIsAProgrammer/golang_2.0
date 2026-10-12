package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// function to connect PostgreSQL database
func Connect(databaseURL string) (*pgxpool.Pool, error) {
	// creating a context without deadlines and no automatic canceallation
	var ctx context.Context = context.Background()

	// parsing database url
	var config *pgxpool.Config
	var err error
	config, err = pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Printf("Unable to parse DATABASE_URL %v", err)
		return nil, err
	}

	// creaeting a connection pool
	var pool *pgxpool.Pool
	pool, err = pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Unable to create connection pool %v", err)
		return nil, err
	}

	// check if PostgreSQL database is reachable
	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping database %v", err)
		pool.Close()
		return nil, err
	}

	log.Println("Successfully connected to the PostgreSQL database!")
	return pool, nil
}
