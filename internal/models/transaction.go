//go:generate mockgen -destination=mock/${GOFILE} -package=${GOPACKAGE}_mock -source=${GOFILE}
package models

import (
	"context"
	"database/sql"
)

type Transaction interface {
	Commit() error
	Rollback() error
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Exec(query string, args ...any) (sql.Result, error)
}
