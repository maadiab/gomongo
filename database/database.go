package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Database interface {
	UserDB
}

type database struct {
	db *sqlx.DB
}

func connect(ctx context.Context) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, "postgres", "user=postgres host=127.0.0.1 dbname=hello password=passwd sslmode=disable")
	return db, err
}

func New(ctx context.Context) (Database, error) {
	db, err := connect(ctx)
	return &database{db: db}, err
}
