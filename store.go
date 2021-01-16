package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Store interface {
	Open() error
	Close() error
}

type dbStore struct {
	db *sqlx.DB
}

var schema = `
CREATE TABLE IF NOT EXISTS auth
(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	access_token TEXT,
	expire_in TEXT,
	refreh_token TEXT
)
`

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "auth.db")
	if err != nil {
		return err
	}
	log.Println("Connected db")
	db.MustExec(schema)

	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}
