package util

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	writer *httptest.ResponseRecorder
	err    error
)

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	writer = httptest.NewRecorder()
	err = errors.New("test handler error")
}

// TestIs test the error type comparability of the HandlerErr struct
func TestIs(t *testing.T) {
	var he HandlerErr
	var err HandlerErr

	// Expect the type of err to match the type of he
	ok := errors.Is(err, he)

	if !ok {
		t.FailNow()
		return
	}
	return
}

// TestErrHandler tests if the error page loads for each case
func TestErrHandler(t *testing.T) {
	fname := "TestErrHandler"

	op := "Initialize template"
	ErrHandler(err, fname, op, time.Now(), writer)
	body := writer.Body.String()
	if strings.Contains(body, "Initialize template") == false {
		t.Error("Error page not loaded")
	}

	op = "Database"
	ErrHandler(err, fname, op, time.Now(), writer)
	body = writer.Body.String()
	if strings.Contains(body, "Database") == false {
		t.Error("Error page not loaded")
	}

	// testing Password case
	op = "Password"
	ErrHandler(err, fname, op, time.Now(), writer)
	body = writer.Body.String()
	if strings.Contains(body, "Incorrect password") == false {
		t.Error("Error page not loaded")
	}

	op = "Session"
	fname = "Logout"
	ErrHandler(err, fname, op, time.Now(), writer)
	body = writer.Body.String()
	if strings.Contains(body, "Session") == false {
		t.Error("Error page not loaded")
	}

}

// test the tmpError function
func TestTmpError(t *testing.T) {
	fname := "template error"
	op := "Initialize template"
	TmpError(err, fname, op, time.Now(), writer)
	if writer.Code != http.StatusInternalServerError {
		t.FailNow()
	}
}

// test the DbError function
func TestUserError(t *testing.T) {
	fname := "UserByEmail"
	op := "Database"
	UserError(err, fname, op, time.Now(), writer)
	if writer.Code != http.StatusBadRequest {
		t.Errorf("\nExpected code %d \t got %d", http.StatusBadRequest, writer.Code)
	}
}

// test the PwError function
func TestPwError(t *testing.T) {
	fname := "CheckPw"
	op := "Password"
	PwError(err, fname, op, time.Now(), writer)
	if writer.Code != http.StatusUnauthorized {
		t.Errorf("\nExpected code %d \t got %d", http.StatusUnauthorized, writer.Code)
	}
}

func TestSessError(t *testing.T) {
	fname := "CreateSession"
	op := "Session"
	SessError(err, fname, op, time.Now(), writer)
	if writer.Code != http.StatusFailedDependency {
		t.Errorf("\nExpected code %d \t got %d", http.StatusFailedDependency, writer.Code)
	}
}
