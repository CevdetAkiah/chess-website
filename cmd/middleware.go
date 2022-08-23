package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf wraps each handler in a CSRFHandler that will call the specified handler if the CSR check succeeds
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
