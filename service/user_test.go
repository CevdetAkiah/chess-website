package service

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"testing"

	_ "github.com/lib/pq"
)

type testOperator User

var err error
var Db *sql.DB

func init() {
	Db, err = sql.Open("postgres", "user=cevdet dbname=website password=cevdet sslmode=disable")
	if err != nil {
		err = fmt.Errorf("Cannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
}
func TestNewService(t *testing.T) {
	to := testOperator{}
	ts := Service{}

	s := NewService(to)

	if reflect.TypeOf(s) != reflect.TypeOf(ts) {
		t.Errorf("Expected type %T, got %T", ts, s)
	}
}

func TestNewUser(t *testing.T) {
	u := User{
		Fname: "Test",
		Lname: "Test",
		Email: "Test",
	}

	to := testOperator{}

	s := NewService(to)

	err = s.NewUser(u)
	if err != nil {
		t.Error(err)
	}

}

func TestDeleteUser(t *testing.T) {
	u := User{
		Fname: "Test",
		Lname: "Test",
		Email: "Test",
	}

	to := testOperator{}

	s := NewService(to)

	err = s.NewUser(u)

	err = s.DeleteUser(u)
	if err != nil {
		t.Error(err)
	}
}

// Methods below are recreated from the database package for the purpose of testing.

// Create inserts the user into the postgres database website table users
func (user testOperator) Create(u User) (err error) {
	statement := "insert into users (fname, lname, email) values ($1, $2, $3) returning id"
	stmnt, err := Db.Prepare(statement)
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
func (user testOperator) Update(u User) (err error) {
	_, err = Db.Exec("update users set email = $1", u.Email)
	if err != nil {
		err = fmt.Errorf("Error updating user: %w", err)
		return
	}
	return
}

// Delete removes a user from the postgres database
func (user testOperator) Delete(u User) (err error) {
	_, err = Db.Exec("delete from users where id = $1", u.Id)
	if err != nil {
		err = fmt.Errorf("Error deleting from users %s, error: %w", u.Fname, err)
		return
	}
	return
}
