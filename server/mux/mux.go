package mux

import (
	"fmt"
	chesswebsocket "go-projects/chess/chesswebsocket"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/route"
	"go-projects/chess/service"
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
	CustomLogger := custom_log.NewLogger()

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
	authSessionHandler, err := route.NewSessionAuthorizer(CustomLogger, DBAccess)
	if err != nil {
		return nil, fmt.Errorf("NewAuthUserHandler error: %b", err)
	}

	gameIDHandler, err := route.NewGameIDRetriever(CustomLogger, DBAccess)
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

	mux.Handle("/ws", websocket.Handler(wsS.HandleWS))
	mux.Get("/healthz", healthzHandler)

	return mux, nil
}
