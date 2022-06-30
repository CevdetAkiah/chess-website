package route

import (
	"go-projects/chess/database/data"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
	"time"
)

func SignupAccount(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	// Set up database service
	r.ParseForm()
	// Get form values
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	pw := r.PostFormValue("password")
	// Encrypt password
	ePw := data.Encrypt(pw)
	// Create user
	u := service.User{
		Name:     name,
		Email:    email,
		Password: ePw,
	}
	// Insert user into database
	err := serve.NewUser(u)
	util.ErrHandler(err, "NewUser", "Database", time.Now(), w)
	http.Redirect(w, r, "/", 302)
}

// Authenticate checks a user exists and creates a session for the user
func Authenticate(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	// Parse the form and get the email
	r.ParseForm()
	email := r.PostFormValue("email")
	// If the user exists, get the user from the database
	user, err := serve.UserByEmail(email)
	util.ErrHandler(err, "UserByEmail", "Database", time.Now(), w)
	// If the password is ok, create a session and set a session cookie
	data.AuthSession(w, r, user, serve)
}
