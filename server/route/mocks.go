package route

import (
	"fmt"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
)

// DB interaction user functions. Abstracts db access
// type DatabaseAccess interface {
// 	CreateUser(user *User) (err error)
// 	Update(user *User) (err error)
// 	DeleteUser(user User) (err error)
// 	UserByEmail(email string) (user User, err error)
// 	CreateSession(user User) (Session, error)
// 	DeleteByUUID(sess Session) (err error)
// 	CheckSession(uuid string) (active bool, err error)
// 	UpdateSession(user User) (err error)
// 	SessionByUuid(uuid string) (Session, error)
// }

type mockDbAccess struct {
	l    custom_log.MagicLogger
	conn service.DatabaseAccess
}

// NewUser stores a new user inside a database
func (db *mockDbAccess) CreateUser(user *service.User) (err error) {
	db.l.Info(fmt.Sprintf("Creating user : %v", user))
	err = db.conn.CreateUser(user)
	return
}

// Update updates a user details in the database
func (db *mockDbAccess) Update(user *service.User) (err error) {
	db.l.Info(fmt.Sprintf("Update user : %v", user))
	err = db.conn.Update(user)
	return
}

// DeleteUser deletes a user from a database
func (db *mockDbAccess) DeleteUser(user service.User) (err error) {
	db.l.Info(fmt.Sprintf("Deleting user : %v", user))
	err = db.conn.DeleteUser(user)
	return
}

func (db *mockDbAccess) UserByEmail(email string) (u service.User, err error) {
	u, err = db.conn.UserByEmail(email)
	return
}

// DeleteByUUID deletes a session from the database using the cookie uuid. Mostly used logging out.
func (db *mockDbAccess) DeleteByUUID(sess service.Session) (err error) {
	err = db.conn.DeleteByUUID(sess)
	return
}

// CreateSession stores a new session in the database on logging in.
func (db *mockDbAccess) CreateSession(u service.User) (sess service.Session, err error) {
	if u.Name == "" {
		return service.Session{}, nil // return empty session
	}
	sess, err = db.conn.CreateSession(u)
	return
}

func (db *mockDbAccess) CheckSession(uuid string) (ok bool, err error) {
	ok, err = db.conn.CheckSession(uuid)
	return
}

func (db *mockDbAccess) SessionByUuid(uuid string) (service.Session, error) {
	session, err := db.conn.SessionByUuid(uuid)
	return session, err
}

func (db *mockDbAccess) UpdateSession(user service.User) (err error) {
	err = db.conn.UpdateSession(user)
	return
}

func (db *mockDbAccess) Print(v string) {
	db.l.Info(v)
}

func (db *mockDbAccess) Printf(format string, a ...any) {
	db.l.Infof(format, a)
}
