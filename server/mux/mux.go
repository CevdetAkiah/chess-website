package mux

import (
	"fmt"
	chesswebsocket "go-projects/chess/chesswebsocket"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/route"
	"go-projects/chess/service"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"golang.org/x/net/websocket"
	// "github.com/go-chi/chi/middleware"
)

var (
	BACKEND_HOST  = os.Getenv("BACKEND_HOST")
	FRONTEND_HOST = os.Getenv("FRONTEND_HOST")
	FRONT_DOMAIN  = os.Getenv("FRONT_DOMAIN")
)

func New(DBAccess service.DatabaseAccess, wsS *chesswebsocket.WsGame) (*chi.Mux, error) {
	mux := chi.NewRouter()

	// mux middleware
	// Nosurf provides each handler with a csrftoken. This provides security against CSRF attacks
	// mux.Use(NoSurf)

	// TODO: look up CSRF protection for chi router
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", BACKEND_HOST, FRONTEND_HOST, FRONT_DOMAIN},
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
	authUserHandler, err := route.NewUserAuthentication(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewAuthUserHandler error: %b", err)
	}

	healthzHandler, err := route.NewHealthz(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewHealthz error: %b", err)
	}

	// Get
	// TODO: user details for profile options
	mux.HandleFunc("/authUser", authUserHandler)
	mux.HandleFunc("/healthz", healthzHandler)

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
