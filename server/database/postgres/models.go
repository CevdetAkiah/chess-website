package postgres

// For connecting to postgres by CLI: sudo -u postgres psql

import (
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	conn *sql.DB
}

func NewDB(pgUser, pgDatabase, pgPassword, pgSSLMode string) *DB {
	conn, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", pgUser, pgDatabase, pgPassword, pgSSLMode))
	if err != nil {
		err = fmt.Errorf("cannot open database with error: %w", err)
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		err = fmt.Errorf("cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	return &DB{
		conn: conn,
	}
}
