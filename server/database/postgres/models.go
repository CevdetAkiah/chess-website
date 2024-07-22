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

func NewDB(config *config.DB) *DB {
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

func (db *DB) SetPermissions(config *config.DB) error {
	query1 := "GRANT ALL PRIVILEGES ON DATABASE $1 TO $2;"
	if _, err := db.conn.Exec(query1, config.PGDatabase, config.PGUser); err != nil {
		return fmt.Errorf("could not grant permissions to user %s on database %s", config.PGUser, config.PGDatabase)
	}
	query2 := "ALTER DATABASE $1 OWNER TO $2;"
	if _, err := db.conn.Exec(query2, config.PGDatabase, config.PGUser); err != nil {
		return fmt.Errorf("could not alter the owner of database %s to user %s", config.PGDatabase, config.PGUser)

	}

	return nil
}
