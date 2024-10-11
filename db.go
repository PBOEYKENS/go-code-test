package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func OpenDatabase() error {
	var err error
	DB, err = sqlx.Open("postgres", "user=pboeykens password=Warlock01 dbname=TestData sslmode=disable")
	if err != nil {
		return err
	}
	return nil
}

func CloseDatabase() error {
	return DB.Close()
}
