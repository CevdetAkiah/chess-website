package service

import (
	"database/sql"
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

type Database interface {
	Create(user User) (err error)
	Update(user User) (err error)
	Delete(user User) (err error)
	CreateSession(user User) (Session, error)
	UserByEmail(email string) (user User, err error)
}

// Server uses interface Storage to CRUD new users
type Server struct {
	Db       *sql.DB
	Operator Database
}

// NewService provides Storage Service where needed
func NewService(operator Database) Server {
	return Server{
		Operator: operator,
	}
}

// NewUser stores a new user inside a database
func (serve *Server) NewUser(user User) (err error) {
	err = serve.Operator.Create(user)
	return
}

func (serve *Server) Update(user User) (err error) {
	err = serve.Operator.Update(user)
	return
}

// DeleteUser deletes a user from a database
func (serve *Server) DeleteUser(u User) (err error) {
	err = serve.Operator.Delete(u)
	return
}

func (serve *Server) CreateSession(u User) (sess Session, err error) {
	sess, err = serve.Operator.CreateSession(u)
	return
}

func (serve *Server) UserByEmail(email string) (u User, err error) {
	u, err = serve.Operator.UserByEmail(email)
	return
}
