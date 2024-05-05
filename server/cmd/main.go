package main

import (
	"context"
	"fmt"
	"go-projects/chess/config"
	postgres "go-projects/chess/database/postgres"
	gameserver "go-projects/chess/gameserver"
	custom_log "go-projects/chess/logger"
	chess_mux "go-projects/chess/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anthdm/hollywood/actor"
)

func main() {
	// config files
	dbConfig := config.NewDBConfig()
	serverConfig := config.NewServerConfig()
	// test database connection
	Db := postgres.NewDB(dbConfig)
	l := custom_log.NewLogger()

	fmt.Println("connected to database chess")

	// handle the chess game server
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		l.Error(err)
	}

	// spawn a new concurrent process for every new ws connection.
	e.Spawn(gameserver.NewGameServer, "server")

	// mux deals with REST api
	mux, err := chess_mux.New(serverConfig, Db)
	if err != nil {
		log.Fatal(err)
	}

	// set up server
	server := &http.Server{
		Addr:         "0.0.0.0:" + serverConfig.Port,
		Handler:      mux,
		WriteTimeout: serverConfig.WriteTimeout,
		ReadTimeout:  serverConfig.ReadTimeout,
	}

	l.Infof("REST api connected to port: %s", serverConfig.Port)

	go func() { // go routine so the enclosed doesn't block
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// sigChan signals when the interrupt or kill signal is received from the OS.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Infof("Received terminate message %v", sig)

	// Graceful shutdown. Users are given 2 minutes to finish their game if the server needs to restart for any reason
	t := time.Now().Add(time.Second * 120)
	tc, _ := context.WithDeadline(context.Background(), t)
	server.Shutdown(tc)
}
