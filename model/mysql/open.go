package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Open is open mysql connection.
func Open(user, dbName string, parseTime bool) (*sql.DB, error) {
	c := &mysql.Config{
		User:      user,
		DBName:    dbName,
		ParseTime: parseTime,
	}

	return sql.Open("mysql", c.FormatDSN())
}
