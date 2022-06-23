package route

import (
	"go-projects/chess/service"
	"net/http"
)

// Request multiplexes http requests
func Request(serv *service.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		switch r.Method {

		// GET retrieves resources
		case "GET":
			if path == "/" {
				Index(w, r, serv)
			} else if path == "/signup" {
				Signup(w, r, serv)
			} else if path == "/errors" {
				ErrorPage(w, r, serv)
			} else if path == "/login" {
				Login(w, r, serv)
			}

			// POST supplies resources
		case "POST":
			if path == "/signupAccount" {
				SignupAccount(w, r, serv)
			} else if path == "/authenticate" {
				Authenticate(w, r, serv)
			}

		}
	}
}
