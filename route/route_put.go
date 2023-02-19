package route

import (
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
