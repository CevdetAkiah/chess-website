package route

import (
	"net/http"
)

// Request multiplexes http requests
func Request(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch r.Method {

	case "GET":
		if path == "/" {
			Index(w, r)
		} else if path == "/signup" {
			Signup(w, r)
		} else if path == "/errors" {
			ErrorPage(w, r)
		} else if path == "/login" {
			Login(w, r)
		}

	case "POST":
		if path == "/signupAccount" {
			SignupAccount(w, r)
		} else if path == "/authenticate" {
			Authenticate(w, r)
		}

	}
}
