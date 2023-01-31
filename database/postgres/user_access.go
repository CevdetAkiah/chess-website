package postgres

import (
	"fmt"
	"go-projects/chess/database/data"
	service "go-projects/chess/service"

	_ "github.com/lib/pq"
)

type UserAccess service.UserDB

// Create inserts the user into the postgres database website table users
func (ua UserAccess) Create(u *service.User) (err error) {
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

// Update alters a user's email in the postgres database
func (ua UserAccess) Update(u *service.User) (err error) {
	_, err = Db.Exec("update users set email = $1 where id = $2", u.Email, u.Id)
	if err != nil {
		err = fmt.Errorf("\nError updating user: %w", err)
		return
	}
	return
}

// Delete removes a user from the postgres database
func (ua UserAccess) Delete(u service.User) (err error) {
	_, err = Db.Exec("delete from users where id = $1", u.Id)
	if err != nil {
		err = fmt.Errorf("\nError deleting from users %s, error: %w", u.Name, err)
		return
	}
	return
}

func (ua UserAccess) UserByEmail(email string) (u service.User, err error) {
	err = Db.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = $1", email).Scan(&u.Id, &u.Uuid, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		err = fmt.Errorf("\nError while getting user by email: %s \n\t Base error: %w", email, err)
	}
	return
}

// TODO: think about re writing the service package so the underlying user type is used like below.

// // Create inserts the user into the postgres database website table users
// func (ua UserAccess) Create(name string, email string, password string, CreatedAt time.Time) (err error) {
// 	statement := "insert into users (uuid, name, email, password, created_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"

// 	stmnt, err := Db.Prepare(statement)
// 	if err != nil {
// 		err = fmt.Errorf("\nError preparing statement to insert user into users table: %w", err)
// 		return
// 	}
// 	defer stmnt.Close()
// 	err = stmnt.QueryRow(data.CreateUUID(), name, email, password, CreatedAt).Scan(&ua.Id, &ua.Uuid, &ua.CreatedAt)
// 	if err != nil {
// 		err = fmt.Errorf("\nError inserting user into users table: %w", err)
// 		return
// 	}

// 	return
// }
