package route

import (
	"fmt"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
)

// SignupAccount is posted from the signup.html template
// SignupAccount creates a user using posted form values and inserts the user into the database
func SignupAccount(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	r.ParseForm()
	// Get form values
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	pw := r.PostFormValue("password")
	// Create user
	user := service.BuildUser(name, email, pw)
	// Insert user into database
	err := DBAccess.NewUser(user)
	if err != nil {
		util.RouteError(w, r, err, "NewUser", "Database")
	}

	http.Redirect(w, r, "/", 200)
}

// Authenticate is activated from the login page
// Authenticate checks a user exists and creates a session for the user
func Authenticate(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	// Parse the form and get the email
	r.ParseForm()
	email := r.PostFormValue("email")
	// If the user exists, get the user from the database
	user, err := DBAccess.UserByEmail(email)

	if err != nil {
		util.RouteError(w, r, err, "UserByEmail", "Database")
	}
	// If password is correct then create session for the user
	if ok := user.Authenticate(r); ok {
		session, err := DBAccess.CreateSession(user)
		if err != nil {
			util.RouteError(w, r, err, "Authenticate handler", "Database")
		}

		session.AssignCookie(w, r)
	}
	// if pw isn't correct then route to error page
	err = fmt.Errorf("Incorrect password")
	util.RouteError(w, r, err, "Authenticate handler", "Password")
}
