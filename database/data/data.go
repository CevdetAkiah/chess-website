package data

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"time"

	"go-projects/chess/service"
	"go-projects/chess/util"

	uuid "github.com/satori/go.uuid"
)

// Encrypt a password
func Encrypt(text string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(text)))
	return
}

// CreateUUID to store in a cookie
func CreateUUID() string {
	sID := uuid.NewV4()
	return sID.String()
}

// AssignCookie puts a cookie into the response writer using the session uuid as the value
func AssignCookie(w http.ResponseWriter, r *http.Request, sess service.Session) {
	cookie := &http.Cookie{
		Name:     "session",
		Value:    sess.Uuid,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", 302)
}

// DeleteCookie removes the cookie from the browser and returns session
func DeleteCookie(w http.ResponseWriter, r *http.Request) (session service.Session) {
	// get the cookie from the request
	cookie, err := r.Cookie("session")
	util.ErrHandler(err, "DeleteCookie", "Session", time.Now(), w)
	// remove cookie from the browser
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)

	// return the session to be removed from the database
	session = service.Session{Uuid: cookie.Value}
	return
}

// AuthSession checks if a users password matches the password for the user in the db
// then creates a session and sets the cookie in the browser
func AuthSession(w http.ResponseWriter, r *http.Request, u service.User, serve service.SessAccess) (err error) {
	if u.Password == Encrypt(r.PostFormValue("password")) {
		session, err := serve.CreateSession(u)
		util.ErrHandler(err, "CreateSession", "Database", time.Now(), w)
		AssignCookie(w, r, session)
	} else {
		err := fmt.Errorf("Bad password")
		util.ErrHandler(err, "Authenticate", "Password", time.Now(), w)
	}
	return
}
