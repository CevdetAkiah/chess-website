package route

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

// CSRFToken - Cross Site Request Forgery Token
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",   // applies to whole site
		Secure:   false, // TODO: in production turn this to true (will be running on https)
		SameSite: http.SameSiteLaxMode,
	})

	
	return csrfHandler
}
