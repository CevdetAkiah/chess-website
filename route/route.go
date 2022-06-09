package route

import (
	"fmt"
	"go-projects/chess/util"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "index")
}

func ErrorPage(w http.ResponseWriter, r *http.Request) {
	util.InitHTML(w, "errors", nil)
}

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from signup")
}
