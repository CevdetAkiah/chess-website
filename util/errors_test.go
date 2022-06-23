package util

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"os"
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

// TODO: test ErrHandler

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
func TestDbError(t *testing.T) {
	fname := "UserByEmail"
	op := "Database"
	DbError(err, fname, op, time.Now(), writer)
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
