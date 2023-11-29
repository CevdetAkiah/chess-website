package main

import (
	"context"
	"fmt"
	chesswebsocket "go-projects/chess/chesswebsocket"
	postgres "go-projects/chess/database/postgres"
	custom_log "go-projects/chess/logger"
	chess_mux "go-projects/chess/mux"
	"go-projects/chess/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	l := custom_log.NewLogger()

	// load env variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	pgUser := os.Getenv("PGUSER")
	pgDatabase := os.Getenv("PGDATABASE")
	pgPassword := os.Getenv("PGPASSWORD")
	pgSSLMode := os.Getenv("PGSSLMODE")
	// test database connection
	Db := postgres.NewDB(pgUser, pgDatabase, pgPassword, pgSSLMode)

	fmt.Println("connected to database chess")

	// set up the database service/access to database
	DBAccess, err := service.NewDBService(
		Db,
		l,
	)
	if err != nil {
		log.Fatal(err)
	}

	chessWebsocket := chesswebsocket.NewWebsocket()

	mux, err := chess_mux.New(DBAccess, chessWebsocket)
	if err != nil {
		log.Fatal(err)
	}

	// set up server
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	DBAccess.Print("Connected to port :8080")

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
	DBAccess.Printf("Received terminate message %v", sig)

	// Graceful shutdown. Users are given 2 minutes to finish their game if the server needs to restart for any reason
	t := time.Now().Add(time.Second * 120)
	tc, _ := context.WithDeadline(context.Background(), t)
	server.Shutdown(tc)
}
