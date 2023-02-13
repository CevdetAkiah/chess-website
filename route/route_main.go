package route

import (
	"go-projects/chess/service"
	"net/http"
)

// Request multiplexes http requests
func Request(DBAccess service.DbService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		switch r.Method {
		// GET retrieves resources
		case "GET":
			if path == "/" {
				Index(w, r, DBAccess)
			} else if path == "/signup" {
				Signup(w, r, DBAccess)
			} else if path == "/errors" {
				ErrorPage(w, r, DBAccess)
			} else if path == "/login" {
				Login(w, r, DBAccess)
			}
			// POST supplies resources to the server
		case "POST":
			if path == "/signupAccount" {
				SignupAccount(w, r, DBAccess)
			} else if path == "/authenticate" {
				Authenticate(w, r, DBAccess)
			} else if path == "/logout" {
				Logout(w, r, DBAccess)
			}

		}
	}
}
