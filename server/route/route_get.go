package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"go-projects/chess/util"
	"net/http"
	"time"
)

// this is used to check the session cookie for log in status each time the client is refreshed
func NewUserAuthentication(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// check for session cookie
		if cookie, err := r.Cookie("session"); err == nil {
			// check if the session cookie is active in the db
			ok, err := DBAccess.CheckSession(cookie.Value)
			if err != nil {
				logger.Error(err)
			}
			// return user data to client if ok
			if ok {
				w.WriteHeader(http.StatusAccepted)
				// processing
				session, err := DBAccess.SessionByUuid(cookie.Value)
				if err != nil {
					logger.Error(err)
				}

				user, err := DBAccess.UserByEmail(session.Email)
				if err != nil {
					logger.Error(err)
				}
				// sending info back
				sendUserDetails(w, user.Name, logger)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}, nil
}

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
