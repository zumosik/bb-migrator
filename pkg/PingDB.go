package pkg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func PingDB(ctx context.Context, driver, uri string) error {
	switch driver {
	case "postgres":
		return pingPostgres(ctx, uri)
	case "mysql":
		return pingMySQL(ctx, uri)
	case "sqlite3":
		return pingSQLite(ctx, uri)
	}

	return fmt.Errorf("unsupported driver: %s", driver)
}

func pingSQLite(ctx context.Context, uri string) error {
	db, err := sql.Open("sqlite3", uri)
	if err != nil {
		return err
	}

	return ping(ctx, db)
}

func pingMySQL(ctx context.Context, uri string) error {
	db, err := sql.Open("mysql", uri)
	if err != nil {
		return err
	}

	return ping(ctx, db)
}

func pingPostgres(ctx context.Context, uri string) error {
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}

	return ping(ctx, db)
}

func ping(ctx context.Context, db *sql.DB) error {
	return db.PingContext(ctx)
}
