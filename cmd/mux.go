package main

import (
	"go-projects/chess/route"
	"go-projects/chess/service"
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/go-chi/chi/middleware"
	"github.com/go-openapi/runtime/middleware"
)

func NewMux(DBAccess service.DbService) *chi.Mux {
	mux := chi.NewRouter()

	// mux middleware
	// Nosurf provides each handler with a csrftoken. This provides security against CSRF attacks
	mux.Use(NoSurf)

	// Pass the request to be handled in the route package
	// Get
	mux.HandleFunc("/", route.Request(DBAccess))
	mux.HandleFunc("/signup", route.Request(DBAccess))
	mux.HandleFunc("/errors", route.Request(DBAccess))
	mux.HandleFunc("/login", route.Request(DBAccess))

	// fileServer serves all static files
	// CSS and JS
	fileServer := http.FileServer(http.Dir("../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	// swagger file
	// TODO: hash the swagger file name for cache busting
	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(options, nil)
	mux.Handle("/docs", sh)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// Post
	mux.HandleFunc("/signupAccount", route.Request(DBAccess))
	mux.HandleFunc("/authenticate", route.Request(DBAccess))
	mux.HandleFunc("/logout", route.Request(DBAccess))

	return mux
}
