package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var pool *pgx.Conn

func DefaultURL() string {
	return os.Getenv("POSTGRES_URL")
}

func GetPool(url string) (*pgx.Conn, error) {
	fmt.Println(url)
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
