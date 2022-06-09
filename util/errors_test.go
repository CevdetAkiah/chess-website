package util

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

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

// TestErrorsHandler test the ErrorsHandler function
func TestErrHandler(t *testing.T) {
	err := errors.New("handler error")
	fname := "test"
	op := "Initialize template"
	w := httptest.NewRecorder()
	ErrHandler(err, fname, op, time.Now(), w)
	if w.Code != http.StatusInternalServerError {
		t.FailNow()
	}
}
