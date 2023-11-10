package route

import (
	"fmt"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
)

// swagger:route POST /signupAccount user createUser
// Send account information to register a new account
// Responses:
//	200:
//		description: "successfully made a new account"
// 		content: application/json

// SignupAccount is posted from the form component of client
// SignupAccount creates a user using posted form values and inserts the user into the database
func signupAccount(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Decode JSON
	userJSON := service.User{}
	err := userJSON.DecodeJSON(r)
	if err != nil {
		DBAccess.Printf("Error while decoding JSON in signupAccount%v", err)
	}
	user := service.BuildUser(userJSON.Name, userJSON.Email, userJSON.Password)

	// Insert user into database
	err = DBAccess.NewUser(user)
	if err != nil {
		// util.RouteError(w, r, err, "NewUser", "Database")
		fmt.Println(err) //TODO: handle this error
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// swagger:route POST /authenticate user authenticateUser
// Send email and password for authentication
// Responses:
//	200:
//		description: "successfully logged in"
// 		content: application/json

// TODO: re write authenticate so api accepts JSON from the frontend
// Authenticate is activated from the login page
// Authenticate checks a user exists and creates a session for the user
func authenticate(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	// Decode JSON
	userJSON := service.User{}
	err := userJSON.DecodeJSON(r)
	if err != nil {
		util.RouteError(w, r, err, "UserByEmail", "Database")
	}
	// If the user exists, get the user from the database
	user, err := DBAccess.UserByEmail(userJSON.Email)
	if err != nil {
		util.RouteError(w, r, err, "UserByEmail", "Database")
	}
	// If password is correct then create session for the user
	if ok := user.Authenticate(userJSON.Password); ok {
		session, err := DBAccess.CreateSession(user)
		if err != nil {
			util.RouteError(w, r, err, "Authenticate handler", "Database")
		}
		session.AssignCookie(w, r)
	} else {
		// if pw isn't correct then route to error page
		err = fmt.Errorf("incorrect password")
		util.RouteError(w, r, err, "Authenticate handler", "Password")
	}

}

// swagger:route POST /logout user logoutUser
// log user out
// Responses:
//	200:
//		description: "successfully logged out"
// 		content: text/html

// logout deletes the session from the browser and database
func logout(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	// send the cookie to be removed from the browser
	session := service.Session{}
	session.DeleteCookie(w, r)
	uuid, err := r.Cookie("session")
	if err != nil {
		util.RouteError(w, r, err, "Logout", "Database")
	}
	session.Uuid = uuid.Value
	// remove the session from the database
	DBAccess.DeleteByUUID(session)
	http.Redirect(w, r, "/", http.StatusFound)
}
