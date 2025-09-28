package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func PostgresDatabase(ctx context.Context, connstring string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25) // max pool
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute) // conn lifetime
	db.SetConnMaxIdleTime(15 * time.Minute) // idle lifetime

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("PostgreSQL database connected")
	return db, nil // return db
}
