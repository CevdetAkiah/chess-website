package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
	"time"
)

func NewGameIDAuthorizer(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}
	return func(w http.ResponseWriter, r *http.Request) {
		// gameCookie is gameID. If no gameCookie, no game is in play.
		if gameCookie, err := r.Cookie("gameID"); err == nil {
			sendUserDetails(w, "", gameCookie.Value, logger)
			return
		}
		fmt.Println("AM I GETTING HERE?")
		w.WriteHeader(http.StatusNoContent)
	}, nil
}

// this is used to check the session cookie for log in status each time the client is refreshed
func NewSessionAuthorizer(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
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
				// processing
				session, err := DBAccess.SessionByUuid(cookie.Value)
				if err != nil {
					logger.Error(err)
				}
				// if the session has timed out remove the session
				if time.Since(session.CreatedAt) > sessionTimeOut {
					err := DBAccess.DeleteByUUID(session)
					if err != nil {
						logger.Error(err)
					}
					err = session.DeleteCookie(w, r)
					if err != nil {
						logger.Error(err)
					}
					w.WriteHeader(http.StatusNoContent)
					return
				}
				w.WriteHeader(http.StatusAccepted)

				user, err := DBAccess.UserByEmail(session.Email)
				if err != nil {
					logger.Error(err)
				}
				// renew session
				DBAccess.UpdateSession(user)
				cookie.MaxAge = session.MaxAge
				http.SetCookie(w, cookie)
				// sending info back
				sendUserDetails(w, user.Name, "", logger)
			} else {
				w.WriteHeader(http.StatusNoContent)
				return
			}
		} else {
			w.WriteHeader(http.StatusNoContent)
			return
		}

	}, nil
}

func NewHealthz(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}, nil

}
