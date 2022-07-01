package service

// NewUser stores a new user inside a database
func (serve DbService) NewUser(user User) (err error) {
	err = serve.UserService.Create(user)
	return
}

func (serve DbService) Update(user User) (err error) {
	err = serve.UserService.Update(user)
	return
}

// DeleteUser deletes a user from a database
func (serve DbService) DeleteUser(u User) (err error) {
	err = serve.UserService.Delete(u)
	return
}

func (serve DbService) UserByEmail(email string) (u User, err error) {
	u, err = serve.UserService.UserByEmail(email)
	return
}
