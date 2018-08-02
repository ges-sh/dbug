package dbugdb

import (
	"context"
	"database/sql"
)

// DB is an interface type which handles all database interferences
type DB interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
