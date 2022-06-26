package postgres

import (
	"database/sql"
	"fmt"
	"go-projects/chess/database/data"
	service "go-projects/chess/service"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

type Operator service.User

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=cevdet dbname=website password=cevdet sslmode=disable")
	if err != nil {
		err = fmt.Errorf("\nCannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
}

// Create inserts the user into the postgres database website table users
func (user Operator) Create(u service.User) (err error) {
	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	stmnt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("\nError preparing statement to insert user into users table: %w", err)
		return
	}
	defer stmnt.Close()
	err = stmnt.QueryRow(data.CreateUUID(), u.Name, u.Email, u.Password, u.CreatedAt).Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError inserting user into users table: %w", err)
		return
	}

	return
}

// Update alters a users email in the postgres database
func (user Operator) Update(u service.User) (err error) {
	_, err = Db.Exec("update users set email = $1", u.Email)
	if err != nil {
		err = fmt.Errorf("\nError updating user: %w", err)
		return
	}
	return
}

// Delete removes a user from the postgres database
func (user Operator) Delete(u service.User) (err error) {
	_, err = Db.Exec("delete from users where id = $1", u.Id)
	if err != nil {
		err = fmt.Errorf("\nError deleting from users %s, error: %w", u.Name, err)
		return
	}
	return
}

func (user Operator) CreateSession(u service.User) (sess service.Session, err error) {
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

func (user Operator) UserByEmail(email string) (u service.User, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError while getting user by email: %s \n\t Base error: %w", email, err)
	}
	return
}

func Retrieve(id int) (u service.User, err error) {
	err = Db.QueryRow("select id, uuid , name, email from users where id = $1", id).Scan(&u.Id, &u.Uuid, &u.Name, &u.Email)
	return
}

// GetAllUsers returns all users from the postgres database
func GetAllUsers() (us []service.User, err error) {
	rows, err := Db.Query("select id, fname, lname, email from users")
	if err != nil {
		err = fmt.Errorf("\nError while getting all users with err: %w", err)
		return
	}
	for rows.Next() {
		user := service.User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email)
		if err != nil {
			err = fmt.Errorf("\nError while retrieving user %s all users with err: %w", user.Name, err)
			return
		}
		us = append(us, user)
	}
	rows.Close()
	return
}
