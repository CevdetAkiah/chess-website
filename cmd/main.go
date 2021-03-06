package main

import (
	"fmt"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/route"
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

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	// set up the database service.
	// Can swap out with any database

	serv := service.DbService{
		Db:             postgres.Db,
		UserService:    postgres.UserAccess{},
		SessionService: postgres.SessionAccess{},
	}

	// Pass the request to be handled in the route package
	mux.HandleFunc("/", route.Request(serv))

	fmt.Println("Connected to port :8080 at", time.Now())
	server.ListenAndServe()

}
