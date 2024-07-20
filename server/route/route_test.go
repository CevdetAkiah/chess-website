package route

import (
	"fmt"
	"go-projects/chess/config"
	"go-projects/chess/database/postgres"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// set up database connection and logger
func setUp() (*postgres.DB, *custom_log.Logger) {
	config := config.NewDB()
	return postgres.NewDB(config), custom_log.NewLogger()
}

// // *
// // GET tests
func TestNewGameIDRetriever(t *testing.T) {
	db, l := setUp()
	mockDbAccess := mockDbAccess{l, db}

	// test request for the case where gameID is present
	reqOK, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	reqOK.AddCookie(&http.Cookie{Name: "gameID", Value: "123"})
	recorderOK := httptest.NewRecorder()

	handlerOK, err := NewGameIDRetriever(l, &mockDbAccess)
	if err != nil {
		t.Fatal(err)
	}
	handlerOK(recorderOK, reqOK)
	assert.Equal(t, http.StatusOK, recorderOK.Code)

	// test request for the case where gameID is not present present
	reqNoID, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorderNoID := httptest.NewRecorder()

	handlerNoID, err := NewGameIDRetriever(l, &mockDbAccess)
	if err != nil {
		t.Fatal(err)
	}
	handlerNoID(recorderNoID, reqNoID)
	assert.Equal(t, http.StatusOK, recorderNoID.Code)
	responseBody := recorderNoID.Body.String()
	if !strings.Contains(responseBody, `"gameID": "new-game"`) {
		t.Error("response body does not contain `gameID: new-game` on no cookie request")
	}
}

func TestNewSessionAuthorizer(t *testing.T) {
	db, l := setUp()
	store := mockDbAccess{l, db}
	req := httptest.NewRequest("GET", "/test", nil)

	testUser := service.NewUser("test", "test@test", "123")

	err := store.CreateUser(testUser)
	if err != nil {
		t.Fatalf("creating user: %b", err)
	}
	// creating a session and checking if it's being renewed
	session, err := store.CreateSession(*testUser)
	if err != nil {
		t.Fatalf("creating session: %b", err)
	}

	req.AddCookie(&http.Cookie{Name: "session", Value: session.Uuid})
	recorder := httptest.NewRecorder()
	handlerRenewSession, err := NewSessionAuthorizer(l, db)
	if err != nil {
		t.Fatalf("creating handler session: %b", err)
	}

	handlerRenewSession(recorder, req)
	assert.Equal(t, http.StatusOK, recorder.Code)
	renewedSession, _ := store.SessionByUuid(session.Uuid)

	if renewedSession.CreatedAt.Sub(session.CreatedAt) <= 0 { // check if session was renewed
		t.Error("session CreatedAt time not refreshed")
	}
	// clean up
	store.DeleteByUUID(session)
	store.DeleteUser(*testUser)

	// * second test * //
	// create a timed out session
	err = store.CreateUser(testUser)
	if err != nil {
		t.Fatalf("creating user: %b", err)
	}
	reqRemoveSession := httptest.NewRequest("GET", "/test", nil)

	// creating a timedout session and checking if it's being removed
	session, err = store.CreateSession(*testUser)
	if err != nil {
		t.Fatalf("creating session: %b", err)
	}
	testUser.CreatedAt = time.Now().Add(-11 * time.Hour)
	store.UpdateSession(*testUser)
	session, _ = store.SessionByUuid(session.Uuid)

	reqRemoveSession.AddCookie(&http.Cookie{Name: "session", Value: session.Uuid})
	recorderRemoveSession := httptest.NewRecorder()
	var handlerRemoveSession func(w http.ResponseWriter, r *http.Request)
	handlerRemoveSession, err = NewSessionAuthorizer(l, db)
	if err != nil {
		t.Fatalf("creating handler remove session: %b", err)
	}
	handlerRemoveSession(recorderRemoveSession, reqRemoveSession)
	assert.Equal(t, http.StatusNoContent, recorderRemoveSession.Code)
	// cleanup
	store.DeleteUser(*testUser)

	//* third test *//
	// test instance where no session is inserted into db
	if err != nil {
		t.Error("creating no session user", err)
	}
	recorderNoSession := httptest.NewRecorder()
	reqNoSession := httptest.NewRequest("GET", "/test", nil)

	handlerNoSession, err := NewSessionAuthorizer(l, db)
	if err != nil {
		t.Error("handlerNoSession: ", err)
	}

	handlerNoSession(recorderNoSession, reqNoSession)
	assert.Equal(t, http.StatusNoContent, recorderNoSession.Result().StatusCode)
}

// *
// POST tests

// test signup a user
func TestNewSignupAccount(t *testing.T) {
	db, l := setUp()
	store := mockDbAccess{l, db}
	recorder := httptest.NewRecorder()

	// set up form values and request
	testUser := service.User{Name: "test", Email: "test@test", Password: "123"}
	payload := fmt.Sprintf(`{"username" : "%s", "email" : "%s", "password" : "%s"}`, testUser.Name, testUser.Email, testUser.Password)

	request := httptest.NewRequest("POST", "/test", strings.NewReader(payload))

	testSignupAccount, err := NewSignupAccount(l, &store)
	if err != nil {
		t.Error("SignupAccount set up:", err)
	}
	testSignupAccount(recorder, request)
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// cleanup
	deleteUser, err := store.UserByEmail("test@test")
	if err != nil {
		t.Error("returning user to clean up: ", err)
	}
	store.DeleteUser(deleteUser)

}

// create a user and test if user can log in
func TestNewLoginHandler(t *testing.T) {
	db, l := setUp()
	store := mockDbAccess{l, db}
	recorder := httptest.NewRecorder()

	testUser := service.User{Name: "test", Email: "test@test", Password: "123"}
	payload := fmt.Sprintf(`{"username" : "%s", "email" : "%s", "password" : "%s"}`, testUser.Name, testUser.Email, testUser.Password)

	request := httptest.NewRequest("GET", "/test", strings.NewReader(payload))
	// create a user and add to DB
	store.CreateUser(service.NewUser(testUser.Name, testUser.Email, testUser.Password))

	testNewLoginHandler, err := NewLoginHandler(l, &store)
	if err != nil {
		t.Error("setting up testNewLoginHandler: ", err)
	}
	testNewLoginHandler(recorder, request)
	assert.Equal(t, http.StatusOK, recorder.Code)
	user, err := store.UserByEmail(testUser.Email)
	if err != nil {
		t.Error("user was not stored in db: ", err)
	}
	session, err := store.SessionByUuid(user.Uuid)
	if err != nil {
		t.Error("session was not created: ", err)
	}

	// cleanup
	store.DeleteByUUID(session)
	store.DeleteUser(user)
}

// *
// DELETE tests

// delete a user from the db
func TestNewDeleteUser(t *testing.T) {
	// set up
	db, l := setUp()
	request := httptest.NewRequest("GET", "/test", nil)
	recorder := httptest.NewRecorder()
	store := mockDbAccess{l, db}
	testUser := service.User{Name: "test", Email: "test@test", Password: "123"}
	store.CreateUser(&testUser)
	session, err := store.CreateSession(testUser)
	if err != nil {
		t.Error("creating session in TestNewDeleteUser: ", err)
	}
	request.AddCookie(&http.Cookie{Name: "session", Value: session.Uuid})

	testNewDeleteUser, err := NewDeleteUser(l, &store)
	if err != nil {
		t.Error("creating NewDeleteUser instance in TestNewDeleteUser: ", err)
	}
	testNewDeleteUser(recorder, request)
	// check status code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// check if session is deleted from the db
	assert.Equal(t, http.StatusOK, recorder.Code)
	ok, _ := store.CheckSession(session.Uuid)
	if ok {
		t.Error("finding session in DB when should be deleted")
	}

	// check if user is deleted from the db
	_, err = store.UserByEmail(testUser.Email)
	if err == nil {
		t.Error("finding user in DB when should be deleted")
	}

}

// log a user out by deleting the session from the db
func TestNewLogoutUser(t *testing.T) {
	// set up
	db, l := setUp()
	request := httptest.NewRequest("GET", "/test", nil)
	recorder := httptest.NewRecorder()
	store := mockDbAccess{l, db}
	testUser := service.User{Name: "test", Email: "test@test", Password: "123"}
	store.CreateUser(&testUser)
	session, err := store.CreateSession(testUser)
	if err != nil {
		t.Error("creating session in TestNewLogoutUser: ", err)
	}
	request.AddCookie(&http.Cookie{Name: "session", Value: session.Uuid})

	testNewLogoutUser, err := NewLogoutUser(l, &store)
	if err != nil {
		t.Error("creating NewLogoutUser Instance in TestNewLogoutUser", err)
	}
	testNewLogoutUser(recorder, request)
	assert.Equal(t, http.StatusNoContent, recorder.Code)
	ok, _ := store.CheckSession(session.Uuid)
	if ok {
		t.Error("TestNewLogoutUser session should be removed from db but not")
	}

	// cleanup
	store.DeleteUser(testUser)

}

// *
// PUT tests
// *
// update a user in the db
func TestNewUpdateUser(t *testing.T) {
	// set up
	db, l := setUp()
	recorder := httptest.NewRecorder()
	store := mockDbAccess{l, db}
	testUser := service.User{Name: "test", Email: "test@test", Password: "123"}
	store.CreateUser(&testUser)
	session, err := store.CreateSession(testUser)
	if err != nil {
		t.Error("creating session in  TestNewUpdateUser: ", err)
	}
	testNewUpdateUser, err := NewUpdateUser(l, db)
	if err != nil {
		t.Error("instantiating NewUpdateUser: ", err)
	}

	// update use email and send details to update the DB
	testUser.Email = "updated@test"
	payload := fmt.Sprintf(`{"username" : "%s", "email" : "%s", "password" : "%s"}`, testUser.Name, testUser.Email, testUser.Password)

	request := httptest.NewRequest("GET", "/test", strings.NewReader(payload))
	request.AddCookie(&http.Cookie{Name: "session", Value: session.Uuid})

	testNewUpdateUser(recorder, request)
	// check behaviour is as expected
	assert.Equal(t, http.StatusOK, recorder.Code)
	updatedUser, err := store.UserByEmail("updated@test")
	if err != nil {
		t.Error("user not updated in TestNewUpdateUser: ", err)
	}

	// cleanup
	store.DeleteByUUID(session)
	store.DeleteUser(updatedUser)
}
