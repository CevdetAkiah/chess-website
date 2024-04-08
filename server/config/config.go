package config

import (
	"os"
	"time"
)

type DBConfig struct {
	PGUser     string
	PGDatabase string
	PGPassword string
	PGSSLMode  string
	Port       string
}

type ServerConfig struct {
	Port         string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		PGUser:     os.Getenv("PGUSER"),
		PGDatabase: os.Getenv("PGDATABASE"),
		PGPassword: os.Getenv("PGPASSWORD"),
		PGSSLMode:  os.Getenv("PGSSLMODE"),
		Port:       os.Getenv("PORT"),
	}
}

func NewServerConfig() *ServerConfig {
	return &ServerConfig{
		Port:         os.Getenv("PORT"),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
