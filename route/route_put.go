package route

import (
	"encoding/json"
	"go-projects/chess/service"
	"net/http"
)

// TODO: update error package to handle PUT errors
// TODO: validation for JSON values

// updateUserName updates a user's username or email in the database
func updateUser(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := decodeUserUpdates(w, r, DBAccess)

	DBAccess.UserService.Update(&user)
}

// updatePassword updates a user's password in the database
func updatePassword(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := decodeUserUpdates(w, r, DBAccess)
	// encrypt new password
	user.Password = service.Encrypt(user.Password)

	DBAccess.UserService.Update(&user)
}

func decodeUserUpdates(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) (user service.User) {
	// get session cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		DBAccess.Printf("can't access cookie with error: %b", err)
	}
	// get the session from db using uuid stored in cookie
	session, err := DBAccess.SessionService.SessionByUuid(cookie.Value)
	if err != nil {
		DBAccess.Printf("Update username error: %v", err)
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
