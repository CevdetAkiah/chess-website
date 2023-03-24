package postgres

// For connecting to postgres by CLI: sudo -u postgres psql

import (
	"database/sql"
	"fmt"
	service "go-projects/chess/service"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=admin dbname=chess password='@ll@long@watchtower1974' sslmode=disable")
	if err != nil {
		err = fmt.Errorf("\nCannot connect to database with error: %w", err)
		log.Fatalln(err)
	}

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
