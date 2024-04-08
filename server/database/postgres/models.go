package postgres

// For connecting to postgres by CLI: sudo -u postgres psql

import (
	"database/sql"
	"fmt"
	"go-projects/chess/config"
	"log"
)

type DB struct {
	conn *sql.DB
}

func NewDB(config *config.DBConfig) *DB {
	conn, err := sql.Open("postgres", fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", config.PGUser, config.PGDatabase, config.PGPassword, config.PGSSLMode))
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
