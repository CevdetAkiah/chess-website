package main

import (
	"fmt"
	postgres "go-projects/chess/database"
	errs "go-projects/chess/errors"
	"go-projects/chess/util"
	"log"
	"net/http"
	"time"
)

type operationError string

var (
	initTemp operationError = "initialize template"
)

func index(w http.ResponseWriter, r *http.Request) {
	err := util.InitHTML(w, "i")
	errs.ErrHandler(err, "index", string(initTemp), time.Now(), w)
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	err := util.InitHTML(w, "errors", nil)
	errs.ErrHandler(err, "errors", string(initTemp), time.Now(), w)
}

func signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from signup")
}

func main() {
	err := postgres.Db.Ping()
	if err != nil {
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	fmt.Println("connected to database website")

	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/errors", errorPage)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}

	fmt.Println("Connected to port :8080 at", time.Now())
	server.ListenAndServe()

}
