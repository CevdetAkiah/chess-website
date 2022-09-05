package util

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/lib/pq"
)

// example of a custom error
// Can use NewError to compare to other errors
var (
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
func ErrHandler(e error, fname string, op string, t time.Time, w http.ResponseWriter) {
	if e != nil {
		switch op {
		case "Initialize template ":
			TmpError(e, fname, op, t, w)
		case "Database":
			UserError(e, fname, op, t, w)
		case "Password":
			PwError(e, fname, op, t, w)
		case "Session":
			SessError(e, fname, op, t, w)
		}
	}
	return
}

// TmpError deals with template errors
func TmpError(e error, fname string, op string, t time.Time, w http.ResponseWriter) {
	var tErr template.ExecError
	errors.As(e, &tErr)
	h := returnHandlerErr(fname, op+tErr.Name, t, e)
	w.WriteHeader(http.StatusInternalServerError)
	InitHTML(w, nil, "errors", false, h.Error())
}

// UserError deals with user database errors
func UserError(e error, fname string, op string, t time.Time, w http.ResponseWriter) {
	var sqlErr *pq.Error
	h := returnHandlerErr(fname, op, t, e)

	// email already exists in database so can't sign up with it
	if errors.As(e, &sqlErr) && sqlErr.Code == pq.ErrorCode(fmt.Sprint(23505)) {
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, nil, "errors", false, dupEmail.Error())
		log.Println(h.Error())

		// Can't find user in database wrong email
	} else if fname == "UserByEmail" {
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, nil, "errors", false, h.Error())
		log.Println(h.Error())

	} else {
		InitHTML(w, nil, "errors", false, h.Error())
		log.Println(h.Error())
	}
}

func SessError(e error, fname string, op string, t time.Time, w http.ResponseWriter) {
	h := returnHandlerErr(fname, op, t, e)

	if fname == "CreateSession" {
		w.WriteHeader(http.StatusFailedDependency)
		InitHTML(w, nil, "errors", false, h.Error())
		log.Println(h.Error())

	} else if fname == "Logout" {
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, nil, "errors", false, h.Error())
		log.Println(h.Error())
	}
}

// PwError deals with password errors
func PwError(e error, fname string, op string, t time.Time, w http.ResponseWriter) {
	h := returnHandlerErr(fname, op, t, e)
	w.WriteHeader(http.StatusUnauthorized)
	InitHTML(w, nil, "errors", false, badpw)
	log.Println(h.Error())
}
