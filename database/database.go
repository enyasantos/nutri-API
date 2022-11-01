package database

import (
	"app/errors"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Connection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/db/foo.db")
	errors.CheckError(err)

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func CloseConnection(db *sql.DB) {
	db.Close()
}
