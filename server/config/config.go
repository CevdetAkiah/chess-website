package config

import (
	"os"
	"time"
)

type DB struct {
	PGUser     string
	PGDatabase string
	PGPassword string
	PGSSLMode  string
	Port       string
}

type Server struct {
	Port               string
	AllowedOrigins     []string
	AllowedMethods     []string
	AllowedHeaders     []string
	ExposedHeaders     []string
	AllowedCredentials bool
	HandlerTimeout     time.Duration
	WriteTimeout       time.Duration
	ReadTimeout        time.Duration
	MaxAge             int
}

func NewDB() *DB {
	return &DB{
		PGUser:     os.Getenv("PGUSER"),
		PGDatabase: os.Getenv("PGDATABASE"),
		PGPassword: os.Getenv("PGPASSWORD"),
		PGSSLMode:  os.Getenv("PGSSLMODE"),
		Port:       os.Getenv("PORT"),
	}
}

func NewServer() Server {
	// declare origins
	localClient := "http://localhost:3000"
	BACKEND_HOST := os.Getenv("BACKEND_HOST")
	FRONTEND_HOST := os.Getenv("FRONTEND_HOST")
	FRONT_DOMAIN := os.Getenv("FRONT_DOMAIN")

	// declare methods
	GET := "GET"
	POST := "POST"
	PUT := "PUT"
	DELETE := "DELETE"
	OPTIONS := "OPTIONS"

	return Server{
		Port:               os.Getenv("PORT"),
		AllowedOrigins:     []string{localClient, BACKEND_HOST, FRONTEND_HOST, FRONT_DOMAIN},
		AllowedMethods:     []string{GET, POST, PUT, DELETE, OPTIONS},
		AllowedHeaders:     []string{"Access-Control-Allow-Origin", "Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowedCredentials: true,
		HandlerTimeout:     200 * time.Millisecond,
		WriteTimeout:       500 * time.Millisecond,
		ReadTimeout:        500 * time.Millisecond,
		MaxAge:             300,
	}
}
