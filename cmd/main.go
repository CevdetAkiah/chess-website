package main

import (
	"fmt"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
	"log"
	"net/http"
	"time"
)

func main() {
	err := postgres.Db.Ping()
	if err != nil {
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	fmt.Println("connected to database website")

	// Set up the database service.
	// Can swap out with any database
	serv := service.DbService{
		Db:             postgres.Db,
		UserService:    postgres.UserAccess{},
		SessionService: postgres.SessionAccess{},
	}

	mux := mux(serv)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	fmt.Println("Connected to port :8080 at", time.Now())

	server.ListenAndServe()
}
