package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

func MySqlDatabase(ctx context.Context, connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(15 * time.Minute)

	if err := db.PingContext(ctx); err != nil {
		defer func() {
			if err := db.Close(); err != nil {
				log.Printf("failed to close db: %v", err)
			}
		}()
		return nil, fmt.Errorf("db.Ping: %w", err)
	}

	return db, nil
}
