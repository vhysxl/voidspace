package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func PostgresDatabase(ctx context.Context, connstring string) (*pgxpool.Pool, error) {
	db, err := pgxpool.ParseConfig(connstring)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.MaxConns = 25
	db.MinConns = 5
	db.MaxConnIdleTime = 30 * time.Minute
	db.MaxConnIdleTime = 15 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("PostgreSQL database connected")
	return pool, nil // return db
}
