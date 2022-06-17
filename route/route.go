package route

import (
	"go-projects/chess/database/data"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "index", nil)
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "errors", nil)
}

// Need to deal with errors and sort out the database so it returns a uuid, and has fields for a password.
// Probably don't need first and last names, just user name.
func Signup(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "signup", nil)
	var id int

	// Set up database service
	po := postgres.Operator{}
	s := service.NewService(po)

	// If data sent to sign up form do the following
	if r.Method == http.MethodPost {
		r.ParseForm()
		// Get form values
		name := r.PostFormValue("Name")
		email := r.PostFormValue("Email")
		pw := r.PostFormValue("Pw")
		// Encrypt password
		ePw := data.Encrypt(pw)
		// Create user
		u := service.User{
			Id:       id,
			Name:     name,
			Email:    email,
			Password: ePw,
		}
		// Insert user into database
		s.NewUser(u)
	}
}
