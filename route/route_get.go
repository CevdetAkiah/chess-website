package route

import (
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
	"time"
)

// swagger:route GET / html Index
// Produce the front page: index.page.html
// Responses:
//	200:
//		description: "successfully loaded the front page"
// 		content: text/html

// Index initialises the index template
func Index(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "index", DBAccess, "")
}

// swagger:route GET /signup html ErrorPage
// Produce the error page: errors.page.html and embeds with the function and operation that caused the error
// Responses:
//	200:
//		description: "successfully loaded the error page"
// 		content: text/html

// ErrorPage initialises the error template
func ErrorPage(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	vals := r.URL.Query()
	util.ErrHandler(w, r, vals.Get("fname"), vals.Get("op"), time.Now())
}

// swagger:route GET /signup html Signup
// Produce the signup page: signup.page.html and allows the user to create a new account
// Responses:
//	200:
//		description: "successfully loaded the signup page"
// 		content: text/html

// Signup initialised the signup template and deals with user registration
func Signup(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "signup", DBAccess, "")
}

// swagger:route GET /login html Login
// Produce the login page: login.page.html and allows the user to log in to the website
// Responses:
//	200:
//		description: "successfully loaded the login page"
// 		content: text/html

// Login initialises the login template
func Login(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "login", DBAccess, "")
}

// TODO: use context to timeout sessions
