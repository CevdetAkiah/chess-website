package service

import (
	"fmt"
	custom_log "go-projects/chess/logger"
)

// DB interaction user functions. Abstracts db access
type DatabaseAccess interface {
	Create(user *User) (err error)
	Update(user *User) (err error)
	Delete(user User) (err error)
	UserByEmail(email string) (user User, err error)
	NewUser(user *User) (err error)
	CreateSession(user User) (Session, error)
	DeleteByUUID(sess Session) (err error)
	CheckSession(uuid string) (active bool, err error)
	UpdateSession(user User) (err error)
	SessionByUuid(uuid string) (Session, error)
}

type DBService struct {
	conn DatabaseAccess
	l    custom_log.MagicLogger
}

// NewDBService creates a new DBService instance with the provided DatabaseAccess implementation
func NewDBService(dba DatabaseAccess, l custom_log.MagicLogger) (*DBService, error) {
	//TODO: check magiclogger != nil, if error then return error
	if l == nil {
		return &DBService{}, fmt.Errorf("NewDBSerice was passed an empty logger interface")
	}

	if dba == nil {
		return &DBService{}, fmt.Errorf("NewDBService was passed an empty database interface")
	}
	return &DBService{conn: dba, l: l}, nil
}

// NewUser stores a new user inside a database
func (db *DBService) NewUser(user *User) (err error) {
	db.l.Info(fmt.Sprintf("Creating user : %v", user))
	err = db.conn.Create(user)
	return
}

// Update updates a user details in the database
func (db *DBService) Update(user *User) (err error) {
	db.l.Info(fmt.Sprintf("Update user : %v", user))
	err = db.conn.Update(user)
	return
}

// DeleteUser deletes a user from a database
func (db *DBService) DeleteUser(user User) (err error) {
	db.l.Info(fmt.Sprintf("Deleting user : %v", user))
	err = db.conn.Delete(user)
	return
}

func (db *DBService) UserByEmail(email string) (u User, err error) {
	u, err = db.conn.UserByEmail(email)
	return
}

// DeleteByUUID deletes a session from the database using the cookie uuid. Mostly used logging out.
func (db *DBService) DeleteByUUID(sess Session) (err error) {
	err = db.conn.DeleteByUUID(sess)
	return
}

// CreateSession stores a new session in the database on logging in.
func (db *DBService) CreateSession(u User) (sess Session, err error) {
	if u.Name == "" {
		return Session{}, nil // return empty session
	}
	sess, err = db.conn.CreateSession(u)
	return
}

func (db *DBService) CheckSession(uuid string) (ok bool, err error) {
	ok, err = db.conn.CheckSession(uuid)
	return
}

func (db *DBService) SessionByUuid(uuid string) (Session, error) {
	session, err := db.conn.SessionByUuid(uuid)
	return session, err
}

func (db *DBService) UpdateSession(user User) (err error) {
	err = db.conn.UpdateSession(user)
	return
}

func (db *DBService) Print(v string) {
	db.l.Info(v)
}

func (db *DBService) Printf(format string, a ...any) {
	db.l.Infof(format, a)
}
