package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
	"time"
)

// swagger:route POST /NewSignupAccount user createUser
// Send account information to register a new account
// Responses:
//	201:
//		description: "successfully made a new account"
// 		content: application/json

// return a handler to add a new user to the database for game state tracking
func NewSignupAccount(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	}

	if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode JSON
		userJSON := service.User{}
		err := userJSON.DecodeJSON(r)
		if err != nil {
			logger.Infof("Error while decoding JSON in signupAccount %v", err)
		}
		user := service.NewUser(userJSON.Name, userJSON.Email, userJSON.Password)

		// Insert user into database
		err = DBAccess.CreateUser(user)
		if err != nil {
			if err.Error() == EMAIL_DUPLICATE {
				logger.Error(err)
				http.Error(w, EMAIL_DUPLICATE, http.StatusConflict)
				return
			} else if err.Error() == USERNAME_DUPLICATE {
				logger.Error(err)
				http.Error(w, USERNAME_DUPLICATE, http.StatusConflict)
				return
			}
		}
		w.WriteHeader(http.StatusCreated)

	}, nil
}

// NewLoginHandler checks a user exists and creates a session for the user so the server can check for state
func NewLoginHandler(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	}
	if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode JSON
		userJSON := service.User{}
		err := userJSON.DecodeJSON(r)
		if err != nil {
			logger.Error(err)
			return
		}
		// If the user exists, get the user from the database
		user, err := DBAccess.UserByEmail(userJSON.Email)
		if err != nil {
			w.Header().Set("WWW-Authenticate", `Basic-realm="Restricted"`)
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}
		user.CreatedAt = time.Now()
		// If password is correct then create session for the user
		ok := user.Authenticate(userJSON.Password)
		if !ok {
			// if pw isn't correct then inform the client
			w.Header().Set("WWW-Authenticate", `Basic-realm="Restricted"`)
			http.Error(w, "Incorrect password", http.StatusUnauthorized)
			return
		}
		session, err := DBAccess.CreateSession(user)
		if err != nil {
			logger.Error(err)
			return
		}
		session.MaxAge = cookieMaxAge
		session.AssignCookie(w, r)

		// send username back to the front end
		sendUserDetails(w, user.Name, "")
	}, nil
}

// swagger:route POST /logout user logoutUser
// log user out
// Responses:
//	200:
//		description: "successfully logged out"
// 		content: text/html
