package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5"
)

var pool *pgx.Conn

var DEFAULT_URL string = os.Getenv("POSTGRES_URL")

func GetPool(url string) (*pgx.Conn, error) {
	if pool != nil {
		return pool, nil
	}

	p, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return nil, err
	}
	pool = p
	return pool, nil
}
