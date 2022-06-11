package route

import (
	"go-projects/chess/util"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "index", nil)
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "errors", nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "signup", nil)
}
