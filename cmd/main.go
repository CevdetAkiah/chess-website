package main

import (
	"fmt"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/route"
	"go-projects/chess/service"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// TODO: add packages used; NoSurf and Chi Router

func main() {
	err := postgres.Db.Ping()
	if err != nil {
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	fmt.Println("connected to database website")

	mux := chi.NewRouter()
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	// Set up the database service.
	// Can swap out with any database
	serv := service.DbService{
		Db:             postgres.Db,
		UserService:    postgres.UserAccess{},
		SessionService: postgres.SessionAccess{},
	}

	// mux middleware

	// Recoverer recovers from panics and provides a stack trace
	mux.Use((middleware.Recoverer))

	// Nosurf provides each handler with a csrftoken. This provides security against CSRF attacks
	mux.Use(NoSurf)

	// Pass the request to be handled in the route package

	// Get
	mux.HandleFunc("/", route.Request(serv))
	mux.HandleFunc("/signup", route.Request(serv))
	mux.HandleFunc("/errors", route.Request(serv))
	mux.HandleFunc("/login", route.Request(serv))
	mux.HandleFunc("/logout", route.Request(serv))

	// Post
	mux.HandleFunc("/signupAccount", route.Request(serv))
	mux.HandleFunc("/authenticate", route.Request(serv))

	fmt.Println("Connected to port :8080 at", time.Now())

	server.ListenAndServe()
}
