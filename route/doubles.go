package route

import (
	"database/sql"
	"go-projects/chess/service"
)

var (
	testServ        service.DbService
	testUserService service.UserAccess
	testSessService service.SessAccess
	testDb              *sql.DB
)
