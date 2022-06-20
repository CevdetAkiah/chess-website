package data

import (
	"crypto/sha1"
	"fmt"
	"net/http"

	"go-projects/chess/service"

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

// SetCookie puts a cookie into the response writer using the session uuid as the value
func SetCookie(w http.ResponseWriter, r *http.Request, sess service.Session) {
	cookie := http.Cookie{
		Name:     "session",
		Value:    sess.Uuid,
		HttpOnly: true,
	}

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", 302)
}
