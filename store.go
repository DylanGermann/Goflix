package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Store interface {
	Open() error
	Close() error
}

type dbStore struct {
	db *sqlx.DB
}

func (store *dbStore) Open() error {
	db, err := sqlx.Connect("sqlite3", "goflix.db")
	if err != nil {
		return err
	}
	log.Println("Connected to DB")
	store.db = db
	return nil
}

func (store *dbStore) Close() error {
	return store.db.Close()
}
