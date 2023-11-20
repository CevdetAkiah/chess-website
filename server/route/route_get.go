package route

import (
	"go-projects/chess/util"
	"net/http"
	"time"
)

// swagger:route GET /signup html ErrorPage
// Produce the error page: errors.page.html and embeds with the function and operation that caused the error
// Responses:
//	200:
//		description: "successfully loaded the error page"
// 		content: text/html

// ErrorPage initialises the error template
func errorPage(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	util.ErrHandler(w, r, vals.Get("fname"), vals.Get("op"), time.Now())
}
