package service

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

func (serve DbService) CheckSession(uuid string) (active bool) {
	active = serve.SessionService.CheckSession(uuid)
	return
}

// TODO: write tests for DeleteByUUID and CreateSession
