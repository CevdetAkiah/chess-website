package route

import (
	"go-projects/chess/service"
	"net/http"
)

func updateUserName(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	session, err := r.Cookie("session")
	if err != nil {
		DBAccess.Printf("can't access session with error: %b", err)
	}
	DBAccess.Printf("Cookie value: %v", session)
}
