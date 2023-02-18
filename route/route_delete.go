package route

import (
	"go-projects/chess/service"
	"net/http"
)

func deleteUser(w http.ResponseWriter, r *http.Request, DBAccess service.DbService) {
	DBAccess.Println("DElETE")
}
