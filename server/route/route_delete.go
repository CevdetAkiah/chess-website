package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
)

// swagger:route DELETE /deleteUser user deleteUser
// Delete user from database and remove session from browser and db
// Responses:
//	200: account delete
//		description: "successfully delete user"
// 		content: application/json

func NewDeleteUser(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger interface is empty")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("database interface is empty")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// get cookie for uuid
		cookie, err := r.Cookie("session")
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// get the session from db using uuid stored in cookie
		session, err := DBAccess.SessionByUuid(cookie.Value)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// get the user from db using the email stored in the session
		user, err := DBAccess.UserByEmail(session.Email)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// delete session from db
		err = DBAccess.DeleteByUUID(session)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// delete user from db
		err = DBAccess.DeleteUser(user)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// remove session from browser cookies
		err = session.DeleteCookie(w, r)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}, nil
}

// logout deletes the session from the browser and database
func NewLogoutUser(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger interface is nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("databaseaccess interface is nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		session := service.Session{}
		cookie, err := r.Cookie("session")

		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		session.Uuid = cookie.Value
		// remove the session from the database and delete the cookie from the browser
		err = session.DeleteCookie(w, r)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// delete the session from the database session table
		err = DBAccess.DeleteByUUID(session)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// report that the request was successful but also no data to be sent back to the client
		w.WriteHeader(http.StatusNoContent)
	}, nil
}
