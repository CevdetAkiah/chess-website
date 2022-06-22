package route

import (
	"go-projects/chess/database/data"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"log"
	"net/http"
	"time"
)

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

// TODO: get login to work, for some reason it's going straight to error page with "users_email_key" error (dup email error)
func Login(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "login.html", nil)
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
	// Need to add function to errHandler to handle session errors
	if err != nil {
		log.Fatalln(err)
	}

	// If the password is ok, create a session and set a session cookie
	if u.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := s.CreateSession(u)
		if err != nil {
			log.Fatalln(err)
		}
		data.SetCookie(w, r, session)

	} else {
		// TODO: change this to "incorrect password" on errors page
		// http.Redirect(w, r, "/login", 302)
		util.ErrHandler(nil, "Authenticate", "Password", time.Now(), w)
	}

}
