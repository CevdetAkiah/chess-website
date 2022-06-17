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
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
}

// Create inserts the user into the postgres database website table users
func (user Operator) Create(u service.User) (err error) {
	statement := "insert into users (name, email, password) values ($1, $2, $3) returning id, uuid, created_at"
	stmnt, err := Db.Prepare(statement)
	if err != nil {
		err = fmt.Errorf("Error preparing statement to insert user into users table: %w", err)
		return
	}
	defer stmnt.Close()
	err = stmnt.QueryRow(data.CreateUUID(), u.Name, u.Email, u.Password).Scan(&u.Id, &u.Uuid, &u.CreatedAt)
	if err != nil {
		err = fmt.Errorf("Error inserting user into users table: %w", err)
		return
	}
	return
}

// Update alters a users email in the postgres database
func (user Operator) Update(u service.User) (err error) {
	_, err = Db.Exec("update users set email = $1", u.Email)
	if err != nil {
		err = fmt.Errorf("Error updating user: %w", err)
		return
	}
	return
}

// Delete removes a user from the postgres database
func (user Operator) Delete(u service.User) (err error) {
	_, err = Db.Exec("delete from users where id = $1", u.Id)
	if err != nil {
		err = fmt.Errorf("Error deleting from users %s, error: %w", u.Name, err)
		return
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
		err = fmt.Errorf("Error while getting all users with err: %w", err)
		return
	}
	for rows.Next() {
		user := service.User{}
		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email)
		if err != nil {
			err = fmt.Errorf("Error while retrieving user %s all users with err: %w", user.Name, err)
			return
		}
		us = append(us, user)
	}
	rows.Close()
	return
}
