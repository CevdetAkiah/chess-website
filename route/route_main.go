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
			switch path {
			case "/":
				Index(w, r, DBAccess)
			case "/signup":
				Signup(w, r, DBAccess)
			case "/errors":
				ErrorPage(w, r, DBAccess)
			case "/login":
				Login(w, r, DBAccess)
			}

			// POST sends resources to the server
		case "POST":
			switch path {
			case "/signupAccount":
				SignupAccount(w, r, DBAccess)
			case "/authenticate":
				Authenticate(w, r, DBAccess)
			case "/logout":
				Logout(w, r, DBAccess)
			}
		}

	}

}
