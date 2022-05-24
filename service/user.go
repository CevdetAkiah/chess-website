package service

type User struct {
	Id    int
	Fname string
	Lname string
	Email string
}

type Storage interface {
	Create(u User) (err error)
	Update(u User) (err error)
	Delete(u User) (err error)
}

// Service uses interface Storage to CRUD new users
type Service struct {
	store Storage
}

// NewService provides Storage Service where needed
func NewService(s Storage) Service {
	return Service{
		store: s,
	}
}

// NewUser stores a new user inside a database
func (s Service) NewUser(u User) (err error) {
	err = s.store.Create(u)
	return
}

// DeleteUser deletes a user from a database
func (s Service) DeleteUser(u User) (err error) {
	err = s.store.Delete(u)
	return
}
