package main

import (
	"context"
	"fmt"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// test database connection
	err := postgres.Db.Ping()
	if err != nil {
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	fmt.Println("connected to database website")

	// set up the database service/access
	DBAccess := service.NewDbService(
		postgres.Db,
		postgres.UserAccess{},
		postgres.SessionAccess{},
		log.New(os.Stdout, "database-api ", log.LstdFlags))

	mux := NewMux(DBAccess)
	// set up server
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	// fmt.Println("Connected to port :8080 at", time.Now())
	DBAccess.Println("Connected to port :8080")
	go func() { // go routine so it doesn't block
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
	DBAccess.Printf("Received terminate message ", sig)

	// Graceful shutdown. Users are given 2 minutes to finish their game if the server needs to restart for any reason
	t := time.Now().Add(time.Second * 120)
	tc, _ := context.WithDeadline(context.Background(), t)
	server.Shutdown(tc)
}
