package db

import (
	"context"
	"database/sql"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

type DB struct {
	db  *sql.DB
	ctx context.Context
}

func New(url string) (*DB, error) {
	db, err := sql.Open("libsql", url)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	return &DB{db, ctx}, nil
}

func (db *DB) Close() {
	db.db.Close()
}

func (db *DB) exec(stmt string, args ...any) (*sql.Result, error) {
	res, err := db.db.ExecContext(db.ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (db *DB) query(stmt string, args ...any) (*sql.Rows, error) {
	res, err := db.db.QueryContext(db.ctx, stmt, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (db *DB) queryRow(stmt string, args ...any) *sql.Row {
	res := db.db.QueryRowContext(db.ctx, stmt, args...)
	return res
}
