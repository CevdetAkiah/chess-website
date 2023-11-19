package route

import (
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
)

// Request multiplexes http requests
func Request(DBAccess *service.DatabaseAccess) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		switch r.Method {
		// POST sends resources to the server
		case "POST":
			switch path {
			case "/signupAccount":
				NewSignupAccount(custom_log.NewLogger())
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
