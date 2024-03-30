package main

import "os"

var (
	// load env variables
	pgUser     = os.Getenv("PGUSER")
	pgDatabase = os.Getenv("PGDATABASE")
	pgPassword = os.Getenv("PGPASSWORD")
	pgSSLMode  = os.Getenv("PGSSLMODE")
	port       = os.Getenv("PORT")
)
