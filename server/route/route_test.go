package route

import (
	"go-projects/chess/database/postgres"
	custom_log "go-projects/chess/logger"
	"go-projects/chess/service"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func setUp() (db *postgres.DB, log *custom_log.Logger) {
	pgUser := os.Getenv("PGUSER")
	pgDatabase := os.Getenv("PGDATABASE")
	pgPassword := os.Getenv("PGPASSWORD")
	pgSSLMode := os.Getenv("PGSSLMODE")
	db = postgres.NewDB(pgUser, pgDatabase, pgPassword, pgSSLMode)
	log = custom_log.NewLogger()

	return
}

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
