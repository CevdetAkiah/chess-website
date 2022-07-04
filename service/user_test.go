package service

import (
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
)

func setUp() {
	mux = http.NewServeMux()
	writer = httptest.NewRecorder()

	mockServ = DbService{
		Db:             testDb,
		UserService:    testUserAccess{},
		SessionService: testSessionAccess{},
	}

	// delete all from testusers table
	// delete all from testsessions table
}

// TODO: learn how to test these functions

func TestNewUser(t *testing.T) {
	user := User{
		Name:     "tom",
		Email:    "tom@email.com",
		Password: "12345",
	}

	err := mockServ.NewUser(&user)
	require.NoError(t, err)
}

// TODO: need to fix Update service as I don't think it's updating a specific user - need to investigate.
func TestUpdate(t *testing.T) {
	user := User{
		Name:     "tom",
		Email:    "tom@email.com",
		Password: "12345",
	}
	err := mockServ.NewUser(&user)

	upDateUser := User{
		Name:     "tom",
		Email:    "tom@email.com",
		Password: "12345",
	}
	err = mockServ.Update(&upDateUser)
	require.NoError(t, err)
}

func TestUserByEmail(t *testing.T) {
	user := User{
		Name:     "tom",
		Email:    "tom@newemail.com",
		Password: "12345",
	}
	err := mockServ.NewUser(&user)

	u, err := mockServ.UserByEmail("tom@newemail.com")
	require.NoError(t, err)
	require.Equal(t, user.Password, u.Password)
}

func TestDeleteUser(t *testing.T) {
	user := User{
		Name:     "tom",
		Email:    "tom@newemail.com",
		Password: "12345",
	}

	err := mockServ.NewUser(&user)

	mockServ.DeleteUser(user)

	if err != nil {
		t.Error("Error: ", err)
	}

	u, err := mockServ.UserByEmail("tom@newemail.com")

	if err == nil {
		t.Errorf("User: %s hasn't been deleted", u.Name)
	}
}
