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
		session.DeleteCookie(w, r)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}, nil
}
