package route

import (
	"go-projects/chess/service"
	"go-projects/chess/util"
	"log"
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
	cookie, err := r.Cookie("session")

	//TODO: add to errHandler for this scenario
	if err != nil {
		log.Fatalln(http.ErrNoCookie)
	}
	session := service.Session{Uuid: cookie.Value}
	serve.DeleteByUUID(session)

	http.Redirect(w, r, "/signup", 302)
}

// TODO: use context to timeout sessions
