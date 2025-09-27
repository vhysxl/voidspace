package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func MySqlDatabase(ctx context.Context, config mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", config.FormatDSN()) // convert config to dsn

	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	} //throw err if nil

	db.SetMaxOpenConns(25) //max pool
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(30 * time.Minute) //conn lifetime
	db.SetConnMaxIdleTime(15 * time.Minute) //idle lifetim

	if err := db.PingContext(ctx); err != nil {
		func() {
			if cerr := db.Close(); cerr != nil {
				log.Printf("failed to close db: %v", cerr)
			}
		}()

		// Ping the database using context. If it fails, close the connection and return the error.
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	log.Println("database connected")
	return db, nil //return db
}
