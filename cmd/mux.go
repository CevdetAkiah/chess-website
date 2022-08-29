package main

import (
	"go-projects/chess/route"
	"go-projects/chess/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func mux(serv service.DbService) *chi.Mux {
	mux := chi.NewRouter()

	// mux middleware
	// Recoverer recovers from panics and provides a stack trace
	mux.Use((middleware.Recoverer))

	// Nosurf provides each handler with a csrftoken. This provides security against CSRF attacks
	mux.Use(NoSurf)

	// Pass the request to be handled in t\he route package

	// fileServer serves all static files
	fileServer := http.FileServer(http.Dir("../static/")) // TODO: change the request header to text/css when css files are served
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	

	// Get
	mux.HandleFunc("/", route.Request(serv))
	mux.HandleFunc("/signup", route.Request(serv))
	mux.HandleFunc("/errors", route.Request(serv))
	mux.HandleFunc("/login", route.Request(serv))
	mux.HandleFunc("/logout", route.Request(serv))

	// Post
	mux.HandleFunc("/signupAccount", route.Request(serv))
	mux.HandleFunc("/authenticate", route.Request(serv))

	return mux
}
