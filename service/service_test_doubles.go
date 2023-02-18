package service

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

type testUserAccess User
type testSessionAccess Session

var (
	testDb *sql.DB
	err    error
)

// Create inserts the user into the postgres database website table users
func (user testUserAccess) Create(u *User) (err error) {
	statement := "insert into testusers (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmnt, err := testDb.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("\nError preparing statement to insert user into testusers table: %w", err)
		return
	}
	defer stmnt.Close()
	err = stmnt.QueryRow(mockCreateUUID(), u.Name, u.Email, u.Password, u.CreatedAt).Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError inserting user into users table: %w", err)
		return
	}

	return
}

// Update alters a users email in the postgres database
func (user testUserAccess) Update(u *User) (err error) {
	_, err = testDb.Exec("update testusers set email = $1 where id = $2", u.Email, u.Id)

	if err != nil {
		err = fmt.Errorf("\nError updating user: %w", err)
		return
	}
	return
}

// Delete removes a user from the postgres database
func (user testUserAccess) Delete(u User) (err error) {
	_, err = testDb.Exec("delete from testusers where id = $1", u.Id)
	if err != nil {
		err = fmt.Errorf("\nError deleting from users %s, error: %w", u.Name, err)
		return
	}
	return
}

func (user testSessionAccess) CreateSession(u User) (sess Session, err error) {
	statement := "insert into testsessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	stmt, err := testDb.Prepare(statement)
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

// Delete session from database
func (session testSessionAccess) DeleteByUUID(sess Session) (err error) {
	statement := "delete from testsessions where uuid = $1"
	stmt, err := testDb.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(sess.Uuid)
	return
}

func (user testUserAccess) UserByEmail(email string) (u User, err error) {
	err = testDb.QueryRow("SELECT id, uuid, name, email, password, created_at FROM testusers WHERE email = $1", email).Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError while getting user by email: %s \n\t Base error: %w", email, err)
	}
	return
}

// CreateUUID to store in a cookie
func mockCreateUUID() string {
	sID := uuid.NewV4()
	return sID.String()
}

// CheckSession checks if the session is active using the given uuid
func (sa testSessionAccess) CheckSession(uuid string) (active bool, err error) {
	err = testDb.QueryRow("SELECT EXISTS(SELECT 1 FROM sessions WHERE uuid = $1)", uuid).Scan(&active)
	if err != nil {
		active = false
		return
	}
	return
}

// SessionByUuid gets session from sessions using given uuid
func (sa testSessionAccess) SessionByUuid(uuid string) (sess Session, err error) {
	err = testDb.QueryRow("SELECT id, email FROM sessions WHERE uuid = $1", uuid).Scan(&sess.Uuid, &sess.Email)
	if err != nil {
		err = fmt.Errorf("\nError getting session by uuid: %w", err)
		return
	}
	return
}
