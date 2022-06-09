package errs

import (
	"errors"
	"testing"
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

// // TestErrorsHandler test the ErrorsHandler function
// func TestErrorsHandler(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/index", main.index)

// 	writer := httptest.NewRecorder()
// 	request, _ := http.NewRequest("GET", "/index", nil)
// 	mux.ServeHTTP(writer, request)

// }
