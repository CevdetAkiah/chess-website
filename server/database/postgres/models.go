package postgres

// For connecting to postgres by CLI: sudo -u postgres psql

import (
	"database/sql"
	"fmt"
	"log"
)

type DB struct {
	conn *sql.DB
}

// TODO: use environemnt variables to replace PW and have the NewDB func accept an env variable as
func NewDB() *DB {
	conn, err := sql.Open("postgres", "user=admin dbname=chess password='@ll@long@watchtower1974' sslmode=disable")
	if err != nil {
		err = fmt.Errorf("cannot open database with error: %w", err)
		log.Fatal(err)
	}
	err = conn.Ping()
	if err != nil {
		err = fmt.Errorf("cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	return &DB{
		conn: conn,
	}
}

// func Retrieve(id int) (u service.User, err error) {
// 	err = Db.QueryRow("select id, uuid , name, email from users where id = $1", id).Scan(&u.Id, &u.Uuid, &u.Name, &u.Email)
// 	return
// }

// // GetAllUsers returns all users from the postgres database
// func GetAllUsers() (us []service.User, err error) {
// 	rows, err := Db.Query("select id, fname, lname, email from users")
// 	if err != nil {
// 		err = fmt.Errorf("\nError while getting all users with err: %w", err)
// 		return
// 	}
// 	for rows.Next() {
// 		user := service.User{}
// 		err = rows.Scan(&user.Id, &user.Uuid, &user.Name, &user.Email)
// 		if err != nil {
// 			err = fmt.Errorf("\nError while retrieving user %s all users with err: %w", user.Name, err)
// 			return
// 		}
// 		us = append(us, user)
// 	}
// 	rows.Close()
// 	return
// }
