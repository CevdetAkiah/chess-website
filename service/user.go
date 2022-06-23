package service

import (
	"time"
)

type User struct {
	// Db *sql.DB
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

type Database interface {
	Create(u User) (err error)
	Update(u User) (err error)
	Delete(u User) (err error)
	CreateSession(u User) (Session, error)
	UserByEmail(email string) (u User, err error)
}

// Service uses interface Storage to CRUD new users
type Service struct {
	operator Database
}

// NewService provides Storage Service where needed
func NewService(d Database) Service {
	return Service{
		operator: d,
	}
}

// NewUser stores a new user inside a database
func (s Service) NewUser(u User) (err error) {
	err = s.operator.Create(u)
	return
}

// DeleteUser deletes a user from a database
func (s Service) DeleteUser(u User) (err error) {
	err = s.operator.Delete(u)
	return
}

func (s Service) CreateSession(u User) (sess Session, err error) {
	sess, err = s.operator.CreateSession(u)
	return
}

func (s Service) UserByEmail(email string) (u User, err error) {
	u, err = s.operator.UserByEmail(email)
	return
}
