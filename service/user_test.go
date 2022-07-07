package service

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	mux    *http.ServeMux
	writer *httptest.ResponseRecorder
)

func TestMain(m *testing.M) {
	setUp()

	code := m.Run()
	os.Exit(code)
}

var (
	mockServ DbService
	user     = User{
		Name:     "tom",
		Email:    "tom@email.com",
		Password: "12345",
	}
)

func setUp() {
	mux = http.NewServeMux()
	writer = httptest.NewRecorder()

	mockServ = DbService{
		Db:             testDb,
		UserService:    testUserAccess{},
		SessionService: testSessionAccess{},
	}
	testDb, err = sql.Open("postgres", "user=cevdet dbname=website password=cevdet sslmode=disable")
	if err != nil {
		err = fmt.Errorf("\nCannot connect to database with error: %w", err)
		log.Fatalln(err)
	}
	query, err := ioutil.ReadFile("../database/testpsql-setup/setup")
	// query,
	if err != nil {
		panic(err)
	}
	if _, err := testDb.Exec(string(query)); err != nil {
		panic(err)
	}

}

func TestNewUser(t *testing.T) {
	err := mockServ.NewUser(&user)
	require.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	err := mockServ.NewUser(&user)

	user.Email = "tom@newemail.com"

	err = mockServ.Update(&user)
	require.NoError(t, err)
	user, err = mockServ.UserByEmail("tom@newemail.com")
	require.NoError(t, err)
}

func TestUserByEmail(t *testing.T) {
	err := mockServ.NewUser(&user)

	u, err := mockServ.UserByEmail("tom@email.com")
	require.NoError(t, err)
	require.Equal(t, user.Password, u.Password)
}

func TestDeleteUser(t *testing.T) {

	err := mockServ.NewUser(&user)

	mockServ.DeleteUser(user)

	if err != nil {
		t.Error("Error: ", err)
	}

	u, err := mockServ.UserByEmail("tom@email.com")

	if err == nil {
		t.Errorf("User: %s hasn't been deleted", u.Name)
	}
}
