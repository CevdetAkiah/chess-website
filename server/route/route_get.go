package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
	"time"
)

// GETs the game id if present or tells the client a new game is being requested
// used to persist the game across re renders
func NewGameIDRetriever(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	}

	if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// gameCookie is gameID. If no gameCookie, no game is in play.
		gameCookie, err := r.Cookie("gameID")
		if err != nil {
			logger.Error(err)
			w.Header().Set("Content-Type", "application/json")
			jsonResponse := `{"gameID": "new-game"}`
			w.Write([]byte(jsonResponse))
			return
		}
		err = sendUserDetails(w, "", gameCookie.Value)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}, nil
}

// this is used to check the session cookie for log in status each time the client is refreshed
func NewSessionAuthorizer(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger was nil")
	}

	if DBAccess == nil {
		return nil, fmt.Errorf("DBA was nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		// check for session cookie
		cookie, err := r.Cookie("session")
		if err != nil {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// check if the session cookie is active in the db
		ok, err := DBAccess.CheckSession(cookie.Value)
		if err != nil {
			logger.Error(err)
		}

		if !ok {
			w.WriteHeader(http.StatusNoContent)
			return
		}

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

		user, err := DBAccess.UserByEmail(session.Email)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// renew session
		err = DBAccess.UpdateSession(user)
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// sending info back
		err = sendUserDetails(w, user.Name, "")
		if err != nil {
			logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		cookie.MaxAge = session.MaxAge
		http.SetCookie(w, cookie)
		// w.WriteHeader(http.StatusAccepted)
	}, nil
}

func NewHealthz() (func(w http.ResponseWriter, r *http.Request), error) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}, nil

}
