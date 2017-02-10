package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Open is open mysql connection.
func Open(user, dbName string) (*sql.DB, error) {
	c := &mysql.Config{
		User:      user,
		DBName:    dbName,
		ParseTime: true,
	}

	return sql.Open("mysql", c.FormatDSN())
}
