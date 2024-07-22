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

	// Get the current working directory
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %v", err)
	}

	// Construct the absolute path to the migrations folder
	path := filepath.Join(workDir, "database", "postgres", "migrations")
	migrationsPath := fmt.Sprintf("file://%s", path)

	// Log the migrations path for debugging
	fmt.Printf("Migrations path: %s\n", migrationsPath)
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
