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

// TODO: Create a database connection to a test user table. Find a way of setting up the test tables when the test starts, and taking them down once the test is finished.

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
	setUp()
	user := User{
		Name:     "tom",
		Email:    "tom@email.com",
		Password: "12345",
	}

	err := mockServ.NewUser(user)
	require.NoError(t, err)

}
