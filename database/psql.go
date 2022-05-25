package postgres

import (
	"database/sql"
	"fmt"
	service "go-projects/chess/service"
	"log"

	_ "github.com/lib/pq"
)

var Psql *sql.DB

type DB service.User

func init() {
	var err error
	Psql, err = sql.Open("postgres", "user=cevdet dbname=website password=cevdet sslmode=disable")
	if err != nil {
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
}

// Create inserts the user into the postgres database website table users
func (user DB) Create(u service.User) (err error) {
	statement := "insert into users (fname, lname, email) values ($1, $2, $3) returning id"
	stmnt, err := Psql.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("Error preparing statement to insert user into users table: %w", err)
		return
	}
	defer stmnt.Close()
	err = stmnt.QueryRow(u.Fname, u.Lname, u.Email).Scan(&u.Id)
	if err != nil {
		err = fmt.Errorf("Error inserting user into users table: %w", err)
		return
	}
	return
}

// Update alters a users email in the postgres database
func (user DB) Update(u service.User) (err error) {
	_, err = Psql.Exec("update users set email = $1", u.Email)
	if err != nil {
		err = fmt.Errorf("Error updating user: %w", err)
		return
	}
	return
}

// Delete removes a user from the postgres database
func (user DB) Delete(u service.User) (err error) {
	_, err = Psql.Exec("delete from users where id = $1", u.Id)
	if err != nil {
		err = fmt.Errorf("Error deleting from users %s, error: %w", u.Fname, err)
		return
	}
	return
}

// GetAllUsers returns all users from the postgres database
func GetAllUsers() (us []service.User, err error) {
	rows, err := Psql.Query("select id, fname, lname, email from users")
	if err != nil {
		err = fmt.Errorf("Error while getting all users with err: %w", err)
		return
	}
	for rows.Next() {
		user := service.User{}
		err = rows.Scan(&user.Id, &user.Fname, &user.Lname, &user.Email)
		if err != nil {
			err = fmt.Errorf("Error while retrieving user %s all users with err: %w", user.Fname, err)
			return
		}
		us = append(us, user)
	}
	rows.Close()
	return
}
