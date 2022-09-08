package route

import (
	"go-projects/chess/database/data"
	"go-projects/chess/service"
	"net/http"
)

var loggedIn bool

// Request multiplexes http requests
func Request(serv service.DbService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// check to see if there is an active session/user is logged in
		loggedIn = data.CheckLogin(r, serv)

		switch r.Method {
		// GET retrieves resources
		case "GET":
			if path == "/" {
				Index(w, r, serv, loggedIn)
			} else if path == "/signup" {
				Signup(w, r, serv, loggedIn)
			} else if path == "/errors" {
				ErrorPage(w, r, serv, loggedIn)
			} else if path == "/login" {
				Login(w, r, serv, loggedIn)
			} else if path == "/logout" {
				Logout(w, r, serv)
			}

			// POST supplies resources to the server
		case "POST":
			if path == "/signupAccount" {
				SignupAccount(w, r, serv)
			} else if path == "/authenticate" {
				Authenticate(w, r, serv)
			}
		}
	}
}
