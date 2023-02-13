package main

import (
	"go-projects/chess/route"
	"go-projects/chess/service"
	"net/http"

	"github.com/gorilla/mux"
)

func NewMux(DBAccess service.DbService) *mux.Router {
	// mux := chi.NewRouter()
	mux := mux.NewRouter()

	// mux middleware
	// Recoverer recovers from panics and provides a stack trace
	// mux.Use((middleware.Recoverer))

	// Nosurf provides each handler with a csrftoken. This provides security against CSRF attacks
	// mux.Use(NoSurf) // TODO: figure out how to get this to work with error handling

	// Pass the request to be handled in the route package
	// Get
	getRouter := mux.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", route.Request(DBAccess))
	getRouter.HandleFunc("/signup", route.Request(DBAccess))
	getRouter.HandleFunc("/errors", route.Request(DBAccess))
	getRouter.HandleFunc("/login", route.Request(DBAccess))
	getRouter.HandleFunc("/logout", route.Request(DBAccess))

	// fileServer serves all static files
	fileServer := http.FileServer(http.Dir("../static/"))
	getRouter.PathPrefix("/static").Handler(http.StripPrefix("/static/", fileServer))

	// Post
	postRouter := mux.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/signupAccount", route.Request(DBAccess))
	postRouter.HandleFunc("/authenticate", route.Request(DBAccess))

	return mux
}
