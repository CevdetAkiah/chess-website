package util

import (
	"errors"
	"fmt"
	"net/http"
	"text/template"
	"time"

	"github.com/go-sql-driver/mysql"
)

// example of a custom error
// Can use NewError to compare to other errors
var HandlerError error

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
	HandlerError = fmt.Errorf("Handler error from handler %s\n \tduring operation %s\n \t\tat time %v\n \t\t\twith base error %w\n", e.Hname, e.Op, e.When, e.Err)
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
			if errors.Is(e, &mysql.MySQLError{}) {
				var sqlErr *mysql.MySQLError
				errors.As(e, &sqlErr)
				// 1062 means aleady exists in database
				if sqlErr.Number == 1062 {
					InitHTML(w, "errors", sqlErr.Message)
				} else {
					h := returnHandlerErr(fname, op, t, e)
					InitHTML(w, "errors", h)
				}
			}
		}
	} else {
		return
	}
}
