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
	badpw        = errors.New("Incorrect email")
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
	fmt.Println("HERE ERROR")
	if e != nil {
		switch op {
		case "Initialize template":
			TmpError(e, fname, op, t, w)
		case "Database":
			DbError(e, fname, op, t, w)
		case "Password":
			PwError(e, fname, op, t, w)
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
	InitHTML(w, "errors", h.Error())
}

// DbError deals with database errors
func DbError(e error, fname string, op string, t time.Time, w http.ResponseWriter) {
	var sqlErr *pq.Error
	fmt.Println("HERE DATABASE ERROR")
	h := returnHandlerErr(fname, op, t, e)

	if errors.As(e, &sqlErr) && sqlErr.Code == pq.ErrorCode(fmt.Sprint(23505)) { // email already exists
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, "errors", dupEmail.Error())
		log.Println(h.Error())

	} else if fname == "UserByEmail" {
		w.WriteHeader(http.StatusBadRequest)
		InitHTML(w, "errors", h.Error())
		log.Println(h.Error())

	} else if fname == "CreateSession" {
		w.WriteHeader(http.StatusFailedDependency)
		InitHTML(w, "errors", h.Error())
		log.Println(h.Error())

	} else {
		InitHTML(w, "errors", h.Error())
	}
}

// PwError deals with password errors
func PwError(e error, fname string, op string, t time.Time, w http.ResponseWriter) {
	fmt.Println("HERE PWERROR")
	h := returnHandlerErr(fname, op, t, e)
	w.WriteHeader(http.StatusUnauthorized)
	InitHTML(w, "errors", badpw)
	log.Println(h.Error())
}
