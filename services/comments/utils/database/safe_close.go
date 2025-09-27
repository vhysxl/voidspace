package utils

import (
	"database/sql"
	"log"
)

func SafeClose(rows *sql.Rows) {
	if err := rows.Close(); err != nil {
		log.Printf("failed to close rows: %v", err)
	}
}
