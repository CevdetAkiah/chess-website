package route

import (
	"encoding/json"
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
)

// decodeUserUpdates retreives JSON from the request and updates the user and session
// TODO http responses to the client for each error; http.StatusForbidden?
func decodeUserUpdates(w http.ResponseWriter, r *http.Request, DBAccess service.DatabaseAccess) (user service.User, err error) {
	// get session cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		return service.User{}, fmt.Errorf("can't access cookie in decodeUserUpdates with error: %b", err)
	}
	// get the session from db using uuid stored in cookie
	session, err := DBAccess.SessionByUuid(cookie.Value)
	if err != nil {
		return service.User{}, fmt.Errorf("can't get session in decodeUserUpdates error: %v", err)
	}

	// get the user from db using the email stored in the session
	user, err = DBAccess.UserByEmail(session.Email)
	if err != nil {
		return service.User{}, fmt.Errorf("get user error in decodeUserUpdates: %v", err)
	}

	// decode request body
	err = user.DecodeJSON(r)
	// TODO: must test the old password before changing the password
	user.Password = service.HashPw(user.Password)
	if err != nil {
		return service.User{}, fmt.Errorf("error while decoding JSON in decodeUserUpdates%v", err)
	}

	// update session
	err = DBAccess.UpdateSession(user)
	if err != nil {
		return service.User{}, fmt.Errorf("error while updating session in decodeUserUpdates%v", err)
	}
	return user, nil
}

func sendUserDetails(w http.ResponseWriter, name string, logger custom_log.MagicLogger) {
	w.Header().Set("Content-Type", "application/json")
	verifiedUser := &service.User{
		Name: name,
	}
	userToSend, err := json.Marshal(verifiedUser)
	if err != nil {
		logger.Error(err)
	}
	w.Write(userToSend)
}
