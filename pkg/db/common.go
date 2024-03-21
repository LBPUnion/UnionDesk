package db

import (
	"context"
	"github.com/jackc/pgx/v5"
)

func getRow[T any](conn *pgx.Conn, sql string, args ...any) (*T, error) {
	rows, err := conn.Query(context.Background(), sql, args...)
	if err != nil {
		return nil, err
	}

	out, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[T])
	return &out, err
}
