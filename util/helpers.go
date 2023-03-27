package util

import (
	"fmt"
	"go-projects/chess/service"
	"net/http"
)

// GetUserName returns the username for the message giver.
// TODO: refactor this into a general function. Doesn't need to be limited to websocket server
func GetUserName(r *http.Request, DBAccess service.DbService) []byte {
	cookie, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err) // TODO: handle this error better
	}
	session, err := DBAccess.SessionService.SessionByUuid(cookie.Value)
	if err != nil {
		fmt.Println(err) // TODO: handle this error better
	}
	user, err := DBAccess.UserService.UserByEmail(session.Email)
	if err != nil {
		fmt.Println(err) // TODO: handle this error better
	}
	username := user.Name + ": "
	return []byte(username)
}

func CheckLogin(r *http.Request, DBAccess service.DbService) (ok bool) {
	cookie, err := r.Cookie("session")
	if err == nil {
		ok, err = DBAccess.SessionService.CheckSession(cookie.Value)
		return
	} else {
		ok = false
		return
	}
}
