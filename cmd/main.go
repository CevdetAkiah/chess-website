package main

import (
	"fmt"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/route"
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

	// Pass the request to be handled in the route package
	mux.HandleFunc("/", route.Request)

	fmt.Println("Connected to port :8080 at", time.Now())
	server.ListenAndServe()

}
