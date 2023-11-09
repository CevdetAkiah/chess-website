package main

import (
	"fmt"
	chesswebsocket "go-projects/chess/chesswebsocket"
	"go-projects/chess/route"
	"go-projects/chess/service"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	// "github.com/go-chi/chi/middleware"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/net/websocket"
)

func NewMux(DBAccess service.DbService, wsS *chesswebsocket.WsGame) *chi.Mux {
	mux := chi.NewRouter()

	// mux middleware
	// Nosurf provides each handler with a csrftoken. This provides security against CSRF attacks
	// mux.Use(NoSurf)

	// TODO: look up CSRF protection for chi router
	mux.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"http://localhost:3000", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

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

	// Put
	mux.HandleFunc("/updateUser", route.Request(DBAccess))
	mux.HandleFunc("/updatePassword", route.Request(DBAccess))

	// Delete
	mux.HandleFunc("/deleteUser", route.Request(DBAccess))

	// Websocket
	mux.Handle("/ws", websocket.Handler(wsS.HandleWS))

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
