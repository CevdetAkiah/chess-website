package route

import (
	"go-projects/chess/database/data"
	postgres "go-projects/chess/database/postgres"
	"go-projects/chess/service"
	"go-projects/chess/util"
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
	var id int

	// Set up database service
	po := postgres.Operator{}
	s := service.NewService(po)

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
	err := s.NewUser(u)
	util.ErrHandler(err, "NewUser", "Database", time.Now(), w)

}
