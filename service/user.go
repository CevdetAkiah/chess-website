package service

import (
	"go-projects/chess/database/data"
	"time"
)

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Storage interface {
	Create(u User) (err error)
	Update(u User) (err error)
	Delete(u User) (err error)
	CreateSession(u User) (data.Session, error)
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

func (s Service) CreateSession(u User) (sess data.Session, err error) {
	sess, err = s.store.CreateSession(u)
	return
}
