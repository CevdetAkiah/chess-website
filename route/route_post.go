package route

import (
	"fmt"
	"go-projects/chess/database/data"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
)

// SignupAccount is posted from the signup.html template
// SignupAccount creates a user using posted form values and inserts the user into the database
func SignupAccount(w http.ResponseWriter, r *http.Request, serve service.DbService) {
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
	err := serve.NewUser(&u)
	if err != nil {
		util.SendError(err)
		url := fmt.Sprintf("/errors?fname=%s&op=%s", "NewUser", "Database")
		http.Redirect(w, r, url, 303)
	}
}

// Authenticate is activated from the login page
// Authenticate checks a user exists and creates a session for the user
func Authenticate(w http.ResponseWriter, r *http.Request, serve service.DbService) {
	// Parse the form and get the email
	r.ParseForm()
	email := r.PostFormValue("email")
	// If the user exists, get the user from the database
	user, err := serve.UserByEmail(email)

	if err != nil {
		util.SendError(err)
		url := fmt.Sprintf("/errors?fname=%s&op=%s", "UserByEmail", "Database")
		http.Redirect(w, r, url, 303)
	}

	// If the password is ok, create a session and set a session cookie
	data.AuthSession(w, r, user, serve)
	return
}
