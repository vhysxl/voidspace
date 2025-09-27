package database

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	"cloud.google.com/go/cloudsqlconn"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"
)

func PostgresDatabase(ctx context.Context, connstring *pgx.ConnConfig, instanceName string) (*sql.DB, error) {
	var opts []cloudsqlconn.Option

	opts = append(opts, cloudsqlconn.WithLazyRefresh())
	d, err := cloudsqlconn.NewDialer(context.Background(), opts...)
	if err != nil {
		return nil, err
	}

	connstring.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		return d.Dial(ctx, instanceName)
	}

	dbURI := stdlib.RegisterConnConfig(connstring)
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %w", err)
	}

	return dbPool, nil // return db
}
