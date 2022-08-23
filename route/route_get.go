package route

import (
	"go-projects/chess/database/data"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
)

// Index initialises the index template
func Index(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	util.InitHTML(w, r, "index", serve)
}

// ErrorPage initialises the error template
func ErrorPage(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	util.InitHTML(w, r, "errors", serve)
}

// Signup initialised the signup template and deals with user registration
func Signup(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	util.InitHTML(w, r, "signup", serve)
}

// Login initialises the login template
func Login(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	util.InitHTML(w, r, "login", serve)
}

func Logout(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	// send the cookie to be removed from the browser and return the session
	session := data.DeleteCookie(w, r)
	// remove the session from the database
	serve.DeleteByUUID(session)
	http.Redirect(w, r, "/", 302)
}

// TODO: use context to timeout sessions
