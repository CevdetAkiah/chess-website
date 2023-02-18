package route

import (
	"go-projects/chess/service"
	"net/http"
)

// TODO: update error package to handle PUT errors
func updateUserName(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	cookie, err := r.Cookie("session")
	if err != nil {
		DBAccess.Printf("can't access session with error: %b", err)
	}
	session, err := DBAccess.SessionService.SessionByUuid(cookie.Value)
	if err != nil {
		DBAccess.Printf("Update username error: %v", err)
	}
	user, err := DBAccess.UserByEmail(session.Email)
	if err != nil {
		DBAccess.Printf("Update username error: %v", err)
	}
	DBAccess.Println(user.Email)
	DBAccess.Println("Update userName")
}

func updateEmail(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	DBAccess.Println("Update email")
}

func updatePassword(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	DBAccess.Println("Update PW")
}
