package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"log"
	"net/http"
)

// TODO: update error package to handle PUT errors
// TODO: validation for JSON values

// swagger:route PUT /updateUser user updateUser
// Update user account email or username
// Responses:
//	200: account updated
//		description: "successfully updated user"
// 		content: application/json

func NewUpdateUser(logger custom_log.MagicLogger, DBAccess service.DatabaseAccess) (func(w http.ResponseWriter, r *http.Request), error) {
	if logger == nil {
		return nil, fmt.Errorf("logger interface is nil")
	} else if DBAccess == nil {
		return nil, fmt.Errorf("database access interface is nil")
	}

	return func(w http.ResponseWriter, r *http.Request) {
		user, err := decodeUserUpdates(w, r, DBAccess)
		if err != nil {
			log.Fatalf("update user error: %b", err)
		}
		DBAccess.Update(&user)
		w.WriteHeader(http.StatusOK)
	}, nil
}
