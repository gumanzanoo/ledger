package postgres

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func New() *pgx.Conn {
	conn, err := pgx.Connect(
		context.Background(), "postgres://user:password@localhost/transactionsdb?sslmode=disable")
	if err != nil {
		panic(err)
	}
	return conn
}
