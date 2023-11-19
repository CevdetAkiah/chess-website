package route

import (
	"go-projects/chess/service"
	"net/http"
)

// TODO: update error package to handle PUT errors
// TODO: validation for JSON values

// swagger:route PUT /updateUser user updateUser
// Update user account email or username
// Responses:
//	200: account updated
//		description: "successfully updated user"
// 		content: application/json

// updateUserName updates a user's username or email in the database
func updateUser(w http.ResponseWriter, r *http.Request, DBAccess *service.DBService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := decodeUserUpdates(w, r, DBAccess)

	DBAccess.Update(&user)
}

// swagger:route PUT /updatePassword user updatePassword
// Update user account password
// Responses:
//	200: account updated
//		description: "successfully updated user password"
// 		content: application/json

// updatePassword updates a user's password in the database
func updatePassword(w http.ResponseWriter, r *http.Request, DBAccess *service.DBService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	user := decodeUserUpdates(w, r, DBAccess)
	// encrypt new password
	user.Password = service.HashPw(user.Password)

	DBAccess.Update(&user)
}
