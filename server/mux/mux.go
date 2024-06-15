package mux

import (
	"fmt"
	"go-projects/chess/config"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/route"
	"go-projects/chess/service"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Multiplexer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

func New(config config.ServerConfig, DBAccess service.DatabaseAccess) (Multiplexer, error) {
	mux := chi.NewRouter()
	CustomLogger := custom_log.NewLogger()

	// mux middleware

	// TODO: look up CSRF protection for chi router
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:   config.AllowedOrigins,
		AllowedMethods:   config.AllowedMethods,
		AllowedHeaders:   config.AllowedHeaders,
		ExposedHeaders:   config.ExposedHeaders,
		AllowCredentials: config.AllowedCredentials,
		MaxAge:           config.MaxAge, // Maximum value not ignored by any of major browsers
	}))
	fmt.Println(config.AllowedOrigins)

	// create handlers
	signupHandler, err := route.NewSignupAccount(config.HandlerTimeout, CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewSignupAccount error: %b", err)
	}
	loginHandler, err := route.NewLoginHandler(config.HandlerTimeout, CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewLoginHandler error: %b", err)
	}

	deleteUserHandler, err := route.NewDeleteUser(config.HandlerTimeout, CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewDeleteUserHandler error: %b", err)
	}
	logoutHandler, err := route.NewLogoutUser(config.HandlerTimeout, CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewLogoutrHandler error: %b", err)
	}
	updateUserHandler, err := route.NewUpdateUser(config.HandlerTimeout, CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewUpdateUserHandler error: %b", err)
	}
	authSessionHandler, err := route.NewSessionAuthorizer(config.HandlerTimeout, CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewAuthUserHandler error: %b", err)
	}

	gameIDHandler, err := route.NewGameIDRetriever(config.HandlerTimeout, CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("newGameIDHandler error: %b", err)
	}

	healthzHandler, err := route.NewHealthz()
	if err != nil {
		return nil, fmt.Errorf("NewHealthz error: %b", err)
	}

	mux.Route("/user", func(r chi.Router) {
		// TODO: get will gather a users details for the profile page
		r.Post("/", signupHandler)       // create a new user
		r.Put("/", updateUserHandler)    // update a user
		r.Delete("/", deleteUserHandler) // delete a user
	})

	mux.Route("/session", func(r chi.Router) {
		r.Get("/", authSessionHandler) // check a session's status
		r.Post("/", loginHandler)      // create a session
		r.Delete("/", logoutHandler)   // delete a session
	})

	mux.Route("/game", func(r chi.Router) {
		r.Get("/", gameIDHandler) // get a game ID. This tells the client if a game is in play
	})

	mux.Get("/healthz", healthzHandler)

	return mux, nil
}
