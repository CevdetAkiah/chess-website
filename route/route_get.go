package route

import (
	"go-projects/chess/database/data"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
)

// Index initialises the index template
func Index(w http.ResponseWriter, r *http.Request, serve *service.DbService) {
	util.InitHTML(w, "index", nil)
}

// ErrorPage initialises the error template
func ErrorPage(w http.ResponseWriter, r *http.Request, serve *service.DbService) {
	util.InitHTML(w, "errors", nil)
}

// Signup initialised the signup template and deals with user registration
func Signup(w http.ResponseWriter, r *http.Request, serve *service.DbService) {
	util.InitHTML(w, "signup", nil)
}

// Login initialises the login template
func Login(w http.ResponseWriter, r *http.Request, serve *service.DbService) {
	util.InitHTML(w, "login", nil)
}

func Logout(w http.ResponseWriter, r *http.Request, serve *service.DbService) {
	data.DeleteSession(w, r, serve)
}

// TODO: use context to timeout sessions
