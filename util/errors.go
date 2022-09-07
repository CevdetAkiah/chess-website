package util

import (
	"errors"
	"fmt"
	"go-projects/chess/service"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/lib/pq"
)

// example of a custom error
// Can use NewError to compare to other errors
var (
	e            error
	HandlerError error
	dupEmail     = errors.New("Email already registered")
	badpw        = errors.New("Incorrect password")
)

// HandlerErr holds
type HandlerErr struct {
	Fname string // function name where error occurred
	Op    string // operation where
	When  time.Time
	Err   error
}

func SendError(sent error) {
	e = sent
}

// returnHandlerErr returns a HandlerErr struct
func returnHandlerErr(name string, operation string, t time.Time, e error) HandlerErr {
	return HandlerErr{
		Fname: name,
		Op:    operation,
		When:  t,
		Err:   e,
	}
}

// Error returns the error for HandlerErr as a string
func (e HandlerErr) Error() string {
	HandlerError = fmt.Errorf("\nError from function %s\n \tduring %s operation \n \t\tat time %v\n \t\t\twith base error %w\n", e.Fname, e.Op, e.When, e.Err)
	return fmt.Sprint(HandlerError)
}

// Is allows comparison of HandlerErr
func (e HandlerErr) Is(other error) bool {
	_, ok := other.(HandlerErr)
	return ok
}

// ErrHandler provides more information for errors that occur in the handlers
func ErrHandler(fname string, op string, t time.Time, w http.ResponseWriter, r *http.Request) {
	fmt.Println("ERRHANDLER: ", fname, op)
	fmt.Println("ERROR: ", e)
	if e != nil {
		switch op {
		case "Initialize template ":
			TmpError(w, r, fname, op, t)
		case "Database":
			UserError(w, r, e, fname, op, t)
		case "Password":
			PwError(w, r, e, fname, op, t)
		case "Session":
			SessError(w, r, e, fname, op, t)
		}
	}
	return
}

// TmpError deals with template errors
func TmpError(w http.ResponseWriter, r *http.Request, fname string, op string, t time.Time) {
	var tErr template.ExecError
	errors.As(e, &tErr)
	h := returnHandlerErr(fname, op+tErr.Name, t, e)
	w.WriteHeader(http.StatusInternalServerError)
	InitHTML(w, r, "errors", false, service.DbService{}, h.Error())
}

// UserError deals with user database errors
func UserError(w http.ResponseWriter, r *http.Request, e error, fname string, op string, t time.Time) {
	var sqlErr *pq.Error
	h := returnHandlerErr(fname, op, t, e)

	// email already exists in database so can't sign up with it
	if errors.As(e, &sqlErr) && sqlErr.Code == pq.ErrorCode(fmt.Sprint(23505)) {
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, r, "errors", false, service.DbService{}, dupEmail.Error())
		log.Println(h.Error())

		// Can't find user in database wrong email
	} else if fname == "UserByEmail" {
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, r, "errors", false, service.DbService{}, h.Error())
		log.Println(h.Error())

	} else {
		InitHTML(w, r, "errors", false, service.DbService{}, h.Error())
		log.Println(h.Error())
	}
}

func SessError(w http.ResponseWriter, r *http.Request, e error, fname string, op string, t time.Time) {
	h := returnHandlerErr(fname, op, t, e)

	if fname == "CreateSession" {
		w.WriteHeader(http.StatusFailedDependency)
		InitHTML(w, r, "errors", false, service.DbService{}, h.Error())
		log.Println(h.Error())

	} else if fname == "Logout" {
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, r, "errors", false, service.DbService{}, h.Error())
		log.Println(h.Error())
	}
}

// PwError deals with password errors
func PwError(w http.ResponseWriter, r *http.Request, e error, fname string, op string, t time.Time) {
	h := returnHandlerErr(fname, op, t, e)
	fmt.Println("PwError: ", e)
	w.WriteHeader(http.StatusUnauthorized)
	InitHTML(w, r, "errors", false, service.DbService{}, badpw.Error())
	log.Println(h.Error())
}

// TODO: change error handling so errors route through the url
