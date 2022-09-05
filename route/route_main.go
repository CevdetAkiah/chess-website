package route

import (
	"fmt"
	"go-projects/chess/service"
	"net/http"
)

// Request multiplexes http requests
func Request(serv service.DbService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		var loggedIn bool

		// check to see if there is an active session/user is logged in
		cookie, _ := r.Cookie("session")
		if cookie != nil {
			loggedIn = serv.SessionService.CheckSession(cookie.Value)
		} else {
			loggedIn = false
		}

		fmt.Println(loggedIn)

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
