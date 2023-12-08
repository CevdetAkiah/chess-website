package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
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
	} else if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

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
			// util.RouteError(w, r, err, "NewUser", "Database")
			fmt.Println("CreateUser error: ", err)
		}
		w.WriteHeader(http.StatusCreated)
	}, nil
}

// NewLoginHandler checks a user exists and creates a session for the user so the server can check for state
func NewLoginHandler(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// Decode JSON
		userJSON := service.User{}
		err := userJSON.DecodeJSON(r)
		if err != nil {
			util.RouteError(w, r, err, "UserByEmail", "Database")
		}
		// If the user exists, get the user from the database
		user, err := DBAccess.UserByEmail(userJSON.Email)
		if err != nil {
			util.RouteError(w, r, err, "UserByEmail", "Database")
		}
		// If password is correct then create session for the user
		if ok := user.Authenticate(userJSON.Password); ok {
			session, err := DBAccess.CreateSession(user)
			if err != nil {
				util.RouteError(w, r, err, "Authenticate handler", "Database")
			}
			session.AssignCookie(w, r)

			// send username back to the front end
			sendUserDetails(w, user.Name, logger)
		} else {
			// if pw isn't correct then route to error page
			err = fmt.Errorf("incorrect password")
			util.RouteError(w, r, err, "Authenticate handler", "Password") //TODO: send the error to the front end to notify user
		}

	}, nil
}

// swagger:route POST /logout user logoutUser
// log user out
// Responses:
//	200:
//		description: "successfully logged out"
// 		content: text/html

// logout deletes the session from the browser and database

func NewLogoutUser(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger interface is nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("databaseaccess interface is nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		session := service.Session{}
		uuid, err := r.Cookie("session")
		if err != nil {
			util.RouteError(w, r, err, "Logout", "Database")
		}
		session.Uuid = uuid.Value
		// remove the session from the database and delete the cookie from the browser
		session.DeleteCookie(w, r)
		DBAccess.DeleteByUUID(session)
		// report that the request was successful but also no data to be sent back to the client
		w.WriteHeader(http.StatusNoContent)
	}, nil
}
