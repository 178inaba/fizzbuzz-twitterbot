package mysql

import (
	"database/sql"

	// This package using mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

// Open is open mysql connection.
func Open(dsn string) (*sql.DB, error) {
	return sql.Open("mysql", dsn)
}
