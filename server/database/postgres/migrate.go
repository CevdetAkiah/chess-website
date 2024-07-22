package postgres

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func (db *DB) InitializeTables() error {
	driver, err := postgres.WithInstance(db.conn, &postgres.Config{})
	if err != nil {
		return err
	}

	// Get the absolute path to the migrations folder
	absPath, err := filepath.Abs("../database/postgres/migrations")
	if err != nil {
		return fmt.Errorf("failed to get absolute path: %v", err)
	}
	migrationsPath := fmt.Sprintf("file://%s", absPath)
	databaseName := os.Getenv("PGDATABASE")
	if databaseName == "" {
		return fmt.Errorf("database name environment variable not set")
	}
	m, err := migrate.NewWithDatabaseInstance(migrationsPath, databaseName, driver)
	if err != nil {
		return fmt.Errorf("NewDatabaseInstance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrating files: %v", err)
	}
	return nil
}
