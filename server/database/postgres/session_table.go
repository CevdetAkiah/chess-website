package postgres

import (
	"fmt"
	service "go-projects/chess/service"
	"time"

	_ "github.com/lib/pq"
)

// type SessionAccess service.DBSession
// type SessionAccess struct

// CreateSession creates a session in the postgres database
func (db *DB) CreateSession(u service.User) (sess service.Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := db.conn.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("\nError preparing statement to create a session for user: %s\n with error: %w ", u.Email, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(u.Uuid, u.Email, u.Id, time.Now()).Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError creating a session for user: %s, %w ", u.Email, err)
		return
	}
	return
}

// Delete session from postgres database
func (db *DB) DeleteByUUID(sess service.Session) (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := db.conn.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("\nError deleting session from database")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(sess.Uuid)
	return
}

// CheckSession checks if the session is active using the given uuid
func (db *DB) CheckSession(uuid string) (active bool, err error) {
	err = db.conn.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE uuid = $1)", uuid).Scan(&active)
	if err != nil {
		active = false
		return
	}
	return
}

// SessionByUuid gets session from sessions using given uuid
func (db *DB) SessionByUuid(uuid string) (sess service.Session, err error) {
	err = db.conn.QueryRow("SELECT uuid, email, created_at FROM sessions WHERE uuid = $1", uuid).Scan(&sess.Uuid, &sess.Email, &sess.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError getting session by uuid: %w", err)
		return
	}
	return
}

func (db *DB) UpdateSession(user service.User) (err error) {
	_, err = db.conn.Exec("UPDATE sessions SET created_at = $1, email = $2 WHERE uuid = $3", user.CreatedAt, user.Email, user.Uuid)
	if err != nil {
		err = fmt.Errorf("\nError updating session by uuid: %w", err)
		return
	}
	return
}
