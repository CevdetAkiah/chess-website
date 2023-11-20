package mux

import (
	"fmt"
	chesswebsocket "go-projects/chess/chesswebsocket"
	custom_log "go-projects/chess/logger"
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
	"golang.org/x/net/websocket"

	// "github.com/go-chi/chi/middleware"
	"github.com/go-openapi/runtime/middleware"
)

func New(DBAccess service.DatabaseAccess, wsS *chesswebsocket.WsGame) (*chi.Mux, error) {
	mux := chi.NewRouter()

	// mux middleware
	// Nosurf provides each handler with a csrftoken. This provides security against CSRF attacks
	// mux.Use(NoSurf)

	// TODO: look up CSRF protection for chi router
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	CustomLogger := custom_log.NewLogger()

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

	// create handlers
	signupHandler, err := route.NewSignupAccount(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewSignupAccount error: %b", err)
	}
	loginHandler, err := route.NewLoginHandler(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewLoginHandler error: %b", err)
	}
	deleteUserHandler, err := route.NewDeleteUser(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewDeleteUserHandler error: %b", err)
	}
	logoutHandler, err := route.NewLogoutUser(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewLogoutrHandler error: %b", err)
	}
	updateUserHandler, err := route.NewUpdateUser(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewUpdateUserHandler error: %b", err)
	}
	// Get
	// TODO: user details for profile options

	// Post
	mux.HandleFunc("/signupAccount", signupHandler)
	mux.HandleFunc("/authenticate", loginHandler)
	mux.HandleFunc("/logout", logoutHandler)

	// // Put
	mux.HandleFunc("/updateUser", updateUserHandler)

	// Delete
	mux.HandleFunc("/deleteUser", deleteUserHandler)

	// Websocket
	mux.Handle("/ws", websocket.Handler(wsS.HandleWS))

	return mux, nil
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
