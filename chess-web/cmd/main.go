package main

import (
	"fmt"
	postgres "go-projects/chess/database"
	"net/http"
	"text/template"

	"log"
)

func index(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("../../templates/index.html"))
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Error while executing index template: ", err)
	}
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
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: nil,
	}
	server.ListenAndServe()

}
