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
//	200: /templates/index.page.html
//		description: "successfully loaded the front page"
// 		content: text/html

// Index initialises the index template
func index(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "index", DBAccess, "")
}

// swagger:route GET /signup html ErrorPage
// Produce the error page: errors.page.html and embeds with the function and operation that caused the error
// Responses:
//	200:
//		description: "successfully loaded the error page"
// 		content: text/html

// ErrorPage initialises the error template
func errorPage(w http.ResponseWriter, r *http.Request) {
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
func signup(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "signup", DBAccess, "")
}

// swagger:route GET /login html Login
// Produce the login page: login.page.html and allows the user to log in to the website
// Responses:
//	200:
//		description: "successfully loaded the login page"
// 		content: text/html

// Login initialises the login template
func login(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "login", DBAccess, "")
}

// swagger:route GET /profile html Profile
// Produce the profile page: profile.page.html and allows the user to log in to the website
// Responses:
//	200:
//		description: "successfully loaded the profile page"
// 		content: text/html

// profile initialises the profile page
func profile(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	util.InitHTML(w, r, "profile", DBAccess, "")
}

// TODO: use context to timeout sessions
