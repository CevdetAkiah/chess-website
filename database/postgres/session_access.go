package postgres

import (
	"fmt"
	service "go-projects/chess/service"

	_ "github.com/lib/pq"
)

type SessionAccess service.DBSession

// CreateSession creates a session in the postgres database
func (sa SessionAccess) CreateSession(u service.User) (sess service.Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("\nError preparing statement to create a session for user: %s\n with error: %w ", u.Email, err)
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(u.Uuid, u.Email, u.Id, u.CreatedAt).Scan(&sess.Id, &sess.Uuid, &sess.Email, &sess.UserId, &sess.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError creating a session for user: %s, %w ", u.Email, err)
		return
	}
	return
}

// Delete session from postgres database
func (sa SessionAccess) DeleteByUUID(sess service.Session) (err error) {
	statement := "delete from sessions where uuid = $1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("\nError deleting session from database")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(sess.Uuid)
	return
}

// CheckSession checks if the session is active using the given uuid
func (sa SessionAccess) CheckSession(uuid string) (active bool, err error) {
	err = Db.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE uuid = $1)", uuid).Scan(&active)
	if err != nil {
		active = false
		return
	}
	return
}

// SessionByUuid gets session from sessions using given uuid
func (sa SessionAccess) SessionByUuid(uuid string) (sess service.Session, err error) {
	err = Db.QueryRow("SELECT uuid, email FROM sessions WHERE uuid = $1", uuid).Scan(&sess.Uuid, &sess.Email)
	if err != nil {
		err = fmt.Errorf("\nError getting session by uuid: %w", err)
		return
	}
	return
}

func (sa SessionAccess) UpdateSession(user service.User) (err error) {
	_, err = Db.Exec("update sessions set email = $1 where uuid = $2", user.Email, user.Uuid)
	if err != nil {
		err = fmt.Errorf("\nError updating session by uuid: %w", err)
		return
	}
	return
}

// SessionById gets session from testsessions using given id
func SessionByUuid(id int) (sess service.Session, err error) {
	err = Db.QueryRow("SELECT id, email FROM sessions WHERE id = $1", id).Scan(&sess.Uuid, &sess.Email)
	if err != nil {
		err = fmt.Errorf("\nError getting session by id: %w", err)
		return
	}
	return
}
