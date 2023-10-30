package service

// handles database interactions for users
type UserDB struct{}

// DB interaction user functions. Abstracts db access
type UserAccess interface {
	Create(user *User) (err error)
	Update(user *User) (err error)
	Delete(user User) (err error)
	UserByEmail(email string) (user User, err error)
}

// NewUser stores a new user inside a database
func (serve DbService) NewUser(user *User) (err error) {
	err = serve.UserService.Create(user)
	return
}

// Update updates a user details in the database
func (serve DbService) Update(user *User) (err error) {
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
