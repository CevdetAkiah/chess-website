package route

import (
	"encoding/json"
	"go-projects/chess/service"
	"net/http"
)

func decodeUserUpdates(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) (user service.User) {
	// get session cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		DBAccess.Printf("can't access cookie in decodeUserUpdates with error: %b", err)
	}
	// get the session from db using uuid stored in cookie
	session, err := DBAccess.SessionService.SessionByUuid(cookie.Value)
	if err != nil {
		DBAccess.Printf("can't get session in decodeUserUpdates error: %v", err)
	}

	// get the user from db using the email stored in the session
	user, err = DBAccess.UserByEmail(session.Email)
	if err != nil {
		DBAccess.Printf("get user error in decodeUserUpdates: %v", err)
	}

	// decode request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		DBAccess.Printf("Error while decoding JSON in decodeUserUpdates%v", err)
	}

	// update session
	err = DBAccess.SessionService.UpdateSession(user)
	if err != nil {
		DBAccess.Printf("Error while updating session in decodeUserUpdates%v", err)
	}
	return user
}
