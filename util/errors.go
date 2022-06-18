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
)

type HandlerErr struct {
	Hname string
	Op    string
	When  time.Time
	Err   error
}

// returnHandlerErr returns a HandlerErr struct
func returnHandlerErr(name string, operation string, t time.Time, e error) HandlerErr {
	return HandlerErr{
		Hname: name,
		Op:    operation,
		When:  t,
		Err:   e,
	}
}

// Error returns the error for HandlerErr as a string
func (e HandlerErr) Error() string {
	HandlerError = fmt.Errorf("\nError from function %s\n \tduring %s operation \n \t\tat time %v\n \t\t\twith base error %w\n", e.Hname, e.Op, e.When, e.Err)
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

		case "Initialize template":
			var tErr template.ExecError
			errors.As(e, &tErr)
			h := returnHandlerErr(fname, op+tErr.Name, t, e)
			w.WriteHeader(http.StatusInternalServerError)
			InitHTML(w, "errors", h.Error())

		case "Database":
			var sqlErr *pq.Error
			h := returnHandlerErr(fname, op, t, e)
			if errors.As(e, &sqlErr) && sqlErr.Code == pq.ErrorCode(fmt.Sprint(23505)) { // email already exists
				w.WriteHeader(http.StatusBadRequest)
				InitHTML(w, "errors", dupEmail.Error())
				log.Println(h.Error())
			} else {
				InitHTML(w, "errors", h.Error())
			}
		}

	}
	return
}
