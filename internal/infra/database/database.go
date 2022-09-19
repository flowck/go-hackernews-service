package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func GetDb(driverName, url string, maxOpenConn int, maxIdleConn int) *sql.DB {
	db, err := sql.Open(driverName, url)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return db
}
