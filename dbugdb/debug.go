package dbugdb

import (
	"context"
	"database/sql"
)

type debug struct {
	db DB
	Logger
}

// New returns new instance of DB with provided *sql.DB and Logger
func New(db DB, logger Logger) DB {
	return &debug{
		db:     db,
		Logger: logger,
	}
}

// ExecContext executes a query without returning any rows. The args are for any placeholder parameters in the query.
func (db *debug) ExecContext(ctx context.Context, query string,
	args ...interface{}) (sql.Result, error) {
	db.Debug(parseQuery(query, args))
	return db.db.ExecContext(ctx, query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT. The args are for any placeholder parameters in the query.
func (db *debug) QueryContext(ctx context.Context, query string,
	args ...interface{}) (*sql.Rows, error) {
	db.Debug(parseQuery(query, args))
	return db.db.QueryContext(ctx, query, args...)
}

// QueryRowContext executes a query that is expected to return at most one row. QueryRowContext always returns a non-nil value. Errors are deferred until Row's Scan method is called. If the query selects no rows, the *Row's Scan will return ErrNoRows. Otherwise, the *Row's Scan scans the first selected row and discards the rest.
func (db *debug) QueryRowContext(ctx context.Context, query string,
	args ...interface{}) *sql.Row {
	db.Debug(parseQuery(query, args))
	return db.db.QueryRowContext(ctx, query, args...)
}
