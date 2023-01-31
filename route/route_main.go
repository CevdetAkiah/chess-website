package route

import (
	"go-projects/chess/database/data"
	"go-projects/chess/service"
	"net/http"
)

var loggedIn bool

// Request multiplexes http requests
func Request(DBAccess service.DbService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// check to see if there is an active session/user is logged in
		loggedIn = data.CheckLogin(r, DBAccess)

		switch r.Method {
		// GET retrieves resources
		case "GET":
			if path == "/" {
				Index(w, r, DBAccess, loggedIn)
			} else if path == "/signup" {
				Signup(w, r, DBAccess, loggedIn)
			} else if path == "/errors" {
				ErrorPage(w, r, DBAccess, loggedIn)
			} else if path == "/login" {
				Login(w, r, DBAccess, loggedIn)
			} else if path == "/logout" {
				Logout(w, r, DBAccess)
			}

			// POST supplies resources to the server
		case "POST":
			if path == "/signupAccount" {
				SignupAccount(w, r, DBAccess)
			} else if path == "/authenticate" {
				Authenticate(w, r, DBAccess)
			}
		}
	}
}





