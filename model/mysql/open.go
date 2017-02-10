package mysql

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// Open is open mysql connection.
func Open(user, addr, dbName string) (*sql.DB, error) {
	c := &mysql.Config{
		User:      user,
		Addr:      addr,
		DBName:    dbName,
		ParseTime: true,
	}

	return sql.Open("mysql", c.FormatDSN())
}
