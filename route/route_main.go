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
				index(w, r, DBAccess)
			case "/signup":
				signup(w, r, DBAccess)
			case "/errors":
				errorPage(w, r)
			case "/login":
				login(w, r, DBAccess)
			case "/profile":
				profile(w, r, DBAccess)
			}

			// POST sends resources to the server
		case "POST":
			switch path {
			case "/signupAccount":
				signupAccount(w, r, DBAccess)
			case "/authenticate":
				authenticate(w, r, DBAccess)
			case "/logout":
				logout(w, r, DBAccess)
			}

			// PUT updates a resource on the server
		case "PUT":
			switch path {
			case "/updateUser":
				updateUser(w, r, DBAccess)
			case "/updatePassword":
				updatePassword(w, r, DBAccess)
			}

		case "DELETE":
			switch path {
			case "/deleteUser":
				deleteUser(w, r, DBAccess)
			}
		}

	}

}
