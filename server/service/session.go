package service

import (
	"net/http"
	"time"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	MaxAge    int
	CreatedAt time.Time
}

// AssignCookie puts a cookie into the response writer using the session uuid as the value
func (s Session) AssignCookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    s.Uuid,
		HttpOnly: true,
		MaxAge:   s.MaxAge,
		SameSite: http.SameSiteNoneMode, // allows cors use http.SameSiteNoneMode in production
	}
	http.SetCookie(w, cookie)

}

// delete the cookie from the browser
func (s Session) DeleteCookie(w http.ResponseWriter, r *http.Request) (err error) {
	// get the cookie from the request
	cookie, err := r.Cookie("session")
	if err != nil {
		return err
	}
	// remove cookie from the browser
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	// return the session to be removed from the database
	s.setUUID(cookie.Value)
	return
}

func (s *Session) setUUID(uuid string) {
	s.Uuid = uuid
}
