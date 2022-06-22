package route

import (
	"go-projects/chess/database/data"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
	"time"
)

// TODO: lookup restful use of http methods and mux

// Index initialises the index template
func Index(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "index", nil)
}

// ErrorPage initialises the error template
func ErrorPage(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "errors", nil)
}

// Signup initialised the signup template and deals with user registration
func Signup(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "signup", nil)
}

func SignupAccount(w http.ResponseWriter, r *http.Request) {
	// Set up database service
	db := postgres.Operator{}
	s := service.NewService(db)

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
	err := s.NewUser(u)
	util.ErrHandler(err, "NewUser", "Database", time.Now(), w)
}

// Login initialises the login template
func Login(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "login", nil)
}

// Authenticate checks a user exists and creates a session for the user
func Authenticate(w http.ResponseWriter, r *http.Request) {
	// Parse the form and get the email
	r.ParseForm()
	email := r.PostFormValue("email")

	// Set up the database service
	db := postgres.Operator{}
	s := service.NewService(db)

	// If the user exists, get the user from the database
	u, err := s.UserByEmail(email)
	util.ErrHandler(err, "UserByEmail", "Database", time.Now(), w)

	// If the password is ok, create a session and set a session cookie
	data.AuthSession(u, s, w, r)
}
