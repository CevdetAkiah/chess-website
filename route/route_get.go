package route

import (
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
	"time"
)

// Index initialises the index template
func Index(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "index", DBAccess, "")
}

// ErrorPage initialises the error template
func ErrorPage(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	vals := r.URL.Query()
	util.ErrHandler(w, r, vals.Get("fname"), vals.Get("op"), time.Now())
}

// Signup initialised the signup template and deals with user registration
func Signup(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "signup", DBAccess, "")
}

// Login initialises the login template
func Login(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "login", DBAccess, "")
}

func Logout(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	// send the cookie to be removed from the browser
	session := service.Session{}
	session.DeleteCookie(w, r)
	// remove the session from the database
	DBAccess.DeleteByUUID(session)
	http.Redirect(w, r, "/", 302)
}

// TODO: use context to timeout sessions
