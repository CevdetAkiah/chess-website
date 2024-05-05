package route

import (
	"context"
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
	"time"
)

// TODO: update error package to handle PUT errors
// TODO: validation for JSON values

// swagger:route PUT /updateUser user updateUser
// Update user account email or username
// Responses:
//	200: account updated
//		description: "successfully updated user"
// 		content: application/json

func NewUpdateUser(HandlerTimeout time.Duration, logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger interface is nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("database access interface is nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), HandlerTimeout)
		defer cancel()

		select {
		case <-ctx.Done():
			logger.Infof("request timeout: %v", ctx.Err())
			w.WriteHeader(http.StatusRequestTimeout)
			return
		default:
			user, err := decodeUserUpdates(w, r, DBAccess)
			if err != nil {
				logger.Error(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			DBAccess.Update(&user)
			w.WriteHeader(http.StatusOK)
		}

	}, nil
}
