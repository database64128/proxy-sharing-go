//go:build cgo

package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func openSQLiteDB(dsn string) (*sql.DB, error) {
	return sql.Open("sqlite3", dsn)
}
