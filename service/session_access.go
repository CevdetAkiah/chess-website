package service

// handles database interactions for sessions
type DBSession struct{}

// DB interaction session functions
type SessAccess interface {
	CreateSession(user User) (Session, error)
	DeleteByUUID(sess Session) (err error)
	CheckSession(uuid string) (active bool, err error)
}

// DeleteByUUID deletes a session from the database using the cookie uuid. Mostly used logging out.
func (serve DbService) DeleteByUUID(sess Session) (err error) {
	err = serve.SessionService.DeleteByUUID(sess)
	return
}

// CreateSession stores a new session in the database on logging in.
func (serve DbService) CreateSession(u User) (sess Session, err error) {
	sess, err = serve.SessionService.CreateSession(u)
	return
}

func (serve DbService) CheckSession(uuid string) (active bool, err error) {
	active, err = serve.SessionService.CheckSession(uuid)
	return
}
