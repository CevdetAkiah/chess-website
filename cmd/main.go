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

	http.HandleFunc("/", route.Index)
	http.HandleFunc("/signup", route.Signup)
	http.HandleFunc("/signupAccount", route.SignupAccount)
	http.HandleFunc("/login", route.SignupAccount)
	http.HandleFunc("/errors", route.ErrorPage)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}

	fmt.Println("Connected to port :8080 at", time.Now())
	server.ListenAndServe()

}
