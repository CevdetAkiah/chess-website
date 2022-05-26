package util

import (
	"fmt"
	postgres "go-projects/chess/database"
	"go-projects/chess/service"
)

// GetAllUsers returns all users from the postgres database
func GetAllUsers() (us []service.User, err error) {
	rows, err := postgres.Psql.Query("select id, fname, lname, email from users")
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
