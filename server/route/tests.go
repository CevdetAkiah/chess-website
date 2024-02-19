package route

import (
	"go-projects/chess/database/postgres"
	custom_log "go-projects/chess/logger"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGameIDAuthorizer(t *testing.T) {
	pgUser := os.Getenv("PGUSER")
	pgDatabase := os.Getenv("PGDATABASE")
	pgPassword := os.Getenv("PGPASSWORD")
	pgSSLMode := os.Getenv("PGSSLMODE")
	db := postgres.NewDB(pgUser, pgDatabase, pgPassword, pgSSLMode)
	l := custom_log.NewLogger()
	mockDbAccess := mockDbAccess{l, db}

	// test request for the case where gameID is present
	reqOK, err := http.NewRequest("GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}

	reqOK.AddCookie(&http.Cookie{Name: "gameID", Value: "123"})
	recorderOK := httptest.NewRecorder()

	handlerOK, err := NewGameIDAuthorizer(l, &mockDbAccess)
	if err != nil {
		t.Fatal(err)
	}
	handlerOK(recorderOK, reqOK)
	assert.Equal(t, http.StatusOK, recorderOK.Code)

	reqNoID, err := http.NewRequest("/GET", "/test", nil)
	if err != nil {
		t.Fatal(err)
	}
	recorderNoID := httptest.NewRecorder()

	handlerNoID, err := NewGameIDAuthorizer(l, &mockDbAccess)
	if err != nil {
		t.Fatal(err)
	}
	handlerNoID(recorderNoID, reqNoID)
	assert.Equal(t, http.StatusOK, recorderNoID.Code)
	assert.Contains(t, recorderNoID.Body.String(), `"gameID":"new-game"`)
}
