package route

import (
	"encoding/json"
	"fmt"
	"go-projects/chess/service"
	"net/http"
)

// TODO: update error package to handle PUT errors

// updateUserName updates a users username in the database
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
	decoder := json.NewDecoder(r.Body)

	decoder.Decode(&user)
	fmt.Println(err)
	DBAccess.UserService.Update(&user)
	DBAccess.Println(user.Name)
	DBAccess.Println("Update userName")
}

// updateEmail updates a users email in the database
func updateEmail(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	DBAccess.Println("Update email")
}

// updatePassword updates a users password in the database
func updatePassword(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	DBAccess.Println("Update PW")
}
