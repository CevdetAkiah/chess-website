package service

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func BuildUser(name, email, password string) *User {
	return &User{Name: name, Email: email, Password: password}
}

func (u *User) Authenticate(r *http.Request) (ok bool) {
	if u.checkPw(r.FormValue("password")) {
		return true
	}
	return false
}

func (u *User) checkPw(formPw string) (ok bool) {
	if bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(formPw)) == nil {
		ok = true
	} else {
		ok = false
	}
	return
}
