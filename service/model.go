package service

import (
	"database/sql"
	"time"
)

type Session struct {
	Id        int
	Uuid      string
	Email     string
	UserId    int
	CreatedAt time.Time
}

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

// DbService uses interface UserAccess to CRUD new users
type DbService struct {
	Db             *sql.DB
	UserService    UserAccess
	SessionService SessAccess
}

// DB interaction user functions. Abstracts db access
// TODO: think about splitting this interface into smaller ones if possible.
type UserAccess interface {
	Create(user *User) (err error)
	Update(user *User) (err error)
	Delete(user User) (err error)
	UserByEmail(email string) (user User, err error)
}

// DB interaction session functions
type SessAccess interface {
	CreateSession(user User) (Session, error)
	DeleteByUUID(sess Session) (err error)
}

// Get from the database
type Retrieval interface {
}
