package main

import (
	"fmt"
	"go-projects/chess/route"
	"go-projects/chess/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

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
	mux.HandleFunc("/profile", route.Request(DBAccess))

	// fileServer serves all static files
	// CSS and JS
	fileServer := http.FileServer(http.Dir("../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	// swagger file
	swaggerFile := hashSwagger()
	options := middleware.RedocOpts{SpecURL: "/" + swaggerFile}
	sh := middleware.Redoc(options, nil)
	mux.Handle("/docs", sh)
	mux.Handle("/"+swaggerFile, http.FileServer(http.Dir("./")))

	// Post
	mux.HandleFunc("/signupAccount", route.Request(DBAccess))
	mux.HandleFunc("/authenticate", route.Request(DBAccess))
	mux.HandleFunc("/logout", route.Request(DBAccess))

	return mux
}

// hash the swagger file name for cache busting reasons
func hashSwagger() string {
	// remove old cache busting files
	filepath.Walk("./", func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			if strings.Contains(f.Name(), "swaggerbust.yaml") {
				os.Remove(f.Name())
			}
		}
		return nil
	})

	// Open swagger yaml file
	original, err := os.Open("swagger.yaml")
	if err != nil {
		panic(err)
	}
	defer original.Close()
	// Make new yammer file
	name := fmt.Sprintf("%d.swaggerbust.yaml", time.Now().Nanosecond())
	new, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer new.Close()
	_, err = io.Copy(new, original)
	if err != nil {
		panic(err)
	}

	return new.Name()
}
