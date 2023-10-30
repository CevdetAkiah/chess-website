package route

import (
	"go-projects/chess/service"
	"net/http"
)

// swagger:route DELETE /deleteUser user deleteUser
// Delete user from database and remove session from browser and db
// Responses:
//	200: account delete
//		description: "successfully delete user"
// 		content: application/json

func deleteUser(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	// get cookie for uuid
	cookie, err := r.Cookie("session")
	if err != nil {
		DBAccess.Printf("can't access cookie in deleteUser with error: %b", err)
	}

	// get the session from db using uuid stored in cookie
	session, err := DBAccess.SessionService.SessionByUuid(cookie.Value)
	if err != nil {
		DBAccess.Printf("can't get session in deleteUser error: %v", err)
	}

	// get the user from db using the email stored in the session
	user, err := DBAccess.UserByEmail(session.Email)
	if err != nil {
		DBAccess.Printf("get user error in deleteUser: %v", err)
	}

	// delete session from db
	err = DBAccess.SessionService.DeleteByUUID(session)
	if err != nil {
		DBAccess.Printf("delete user from db error %b", err)
	}

	// delete user from db
	err = DBAccess.UserService.Delete(user)
	if err != nil {
		DBAccess.Printf("delete user from db error %b", err)
	}

	// remove session from browser cookies
	session.DeleteCookie(w, r)
}
